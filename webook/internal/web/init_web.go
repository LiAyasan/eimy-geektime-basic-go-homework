package web

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	registerUsersRoutes(server)
	return server
}

// registerUsersRoutes 第二种注册路由的写法 统一写在一个文件里
func registerUsersRoutes(server *gin.Engine) {
	//// 注册路由都写在 main.go 会显得冗余
	//u := &UserHandler{}
	//// 注册
	//server.POST("/users/signup", u.SingUp)
	//// 登录
	//server.POST("/users/login", u.Login)
	//// 编辑
	//server.POST("/users/edit", u.Edit)
	//// 查看
	//server.GET("/users/profile", u.Profile)
}
