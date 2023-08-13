package main

import (
	"eimy-geektime-basic-go-homework/webook/internal/repository"
	"eimy-geektime-basic-go-homework/webook/internal/repository/dao"
	"eimy-geektime-basic-go-homework/webook/internal/service"
	"eimy-geektime-basic-go-homework/webook/internal/web"
	"eimy-geektime-basic-go-homework/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer()

	u := initUser(db)
	u.RegisterRoutes(server)
	server.Run(":8080")
}

func initWebServer() *gin.Engine {

	server := gin.Default()

	server.Use(func(context *gin.Context) {
		println("第一个middleware")
	})

	// 处理跨域
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"https://foo.com", "http://localhost:3000"},
		// 默认是任何方法
		//AllowMethods:  []string{"PUT", "PATCH", "POST"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},

		// 是否允许你带 cookie 之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// 如果 origin 包含 以下，就通过
			if strings.Contains(origin, "http://localhost") {

				// 你的开发环境
				return true
			}
			// 线上环境
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	// 登录1 启动session
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))
	// 登录校验
	// 登录3 检查session
	server.Use(middleware.NewLoginMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())
	return server
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		// 我只会在初始化过程 panic
		panic(err)
	}

	db = db.Debug()

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}
