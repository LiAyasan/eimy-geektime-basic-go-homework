package web

import (
	"eimy-geektime-basic-go-homework/webook/internal/domain"
	"eimy-geektime-basic-go-homework/webook/internal/service"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"unicode/utf8"
)

// UserHandler 定义跟 User 有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		// 标准正则库不支持下面的写法
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,72}$`
	)
	return &UserHandler{
		svc:         svc,
		emailExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (u *UserHandler) RegisterRoutesV1(server *gin.RouterGroup) {
	// 直接传入分组
	server.POST("/signup", u.SingUp)
	server.POST("/login", u.Login)
	server.POST("/edit", u.Edit)
	server.GET("/profile", u.Profile)
}

// RegisterRoutes 第三种注册路由的写法 分散在各自功能实现的文件里
func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	// 分组注册
	//ug := server.Group("/users")
	//ug.GET("/profile", u.Profile)

	// 注册
	server.POST("/users/signup", u.SingUp)
	// 登录
	server.POST("/users/login", u.Login)
	// 编辑
	server.POST("/users/edit", u.Edit)
	// 查看
	server.GET("/users/profile", u.Profile)
}

func (u *UserHandler) SingUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	// 获取数据
	var req SignUpReq
	// Bind() 根据 Content-Type 来解析数据到 req 里面
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 正则校验
	// 预编译
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}

	if req.ConfirmPassword != req.Password {
		fmt.Println(req.ConfirmPassword, req.Password)
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	// 调用 svc 的方法
	// ctx 也可传 ctx.Request().Context
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	// 数据入库
}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 校验登录态

	// 尝试登录
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名或密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	// 登录成功了
	// 设置 session
	// 登录2 保存session
	sess := sessions.Default(ctx)
	// 你要放在 session 里面的值
	sess.Set("userId", user.Id)
	sess.Save()

	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		Nickname string `json:"nickname"`
		Birthday string `json:"birthday"`
		Details  string `json:"details"`
	}

	// 获取请求字段
	var req EditReq
	if err := ctx.Bind(&req); err != nil {
		ctx.String(http.StatusOK, "系统故障")
		return
	}

	// 校验字段是否合法
	nameLen := utf8.RuneCountInString(req.Nickname)
	if nameLen < 0 || nameLen > 8 {
		ctx.String(http.StatusOK, "昵称长度不合法")
		return
	}

	// 允许传空
	if len(req.Birthday) != 0 {
		if _, err := time.Parse("2006-01-02", req.Birthday); err != nil {
			ctx.String(http.StatusOK, "生日日期有误")
			return
		}
	}

	detailsLen := utf8.RuneCountInString(req.Details)
	if detailsLen < 0 || detailsLen > 300 {
		ctx.String(http.StatusOK, "简介长度不合法")
		return
	}

	// 获取cookie的用户id
	sess := sessions.Default(ctx)
	userId := sess.Get("userId").(int64)
	err := u.svc.Edit(ctx, userId, req.Nickname, req.Birthday, req.Details)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.String(http.StatusOK, "修改成功")
	return
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	// 获取cookie的用户id
	sess := sessions.Default(ctx)
	userId := sess.Get("userId").(int64)

	user, err := u.svc.Profile(ctx, userId)
	if err != nil {
		ctx.String(http.StatusOK, "系统故障")
		return
	}

	ctx.String(http.StatusOK, user)
	return
}
