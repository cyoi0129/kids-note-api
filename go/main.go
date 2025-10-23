package main

import (
	"kids-note-api/controllers"
	"kids-note-api/models"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("TZ", "UTC") // タイムゾーンの設定
	r := gin.Default()

	r.Use(cors.New(cors.Config{

		AllowOrigins: []string{ // アクセスを許可したいアクセス元
			"http://localhost:3000",
		},

		AllowMethods: []string{ // アクセスを許可したいHTTPメソッド
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},

		AllowHeaders: []string{ // 許可したいHTTPリクエストヘッダ
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},

		AllowCredentials: true, // cookieなどの情報を必要とするかどうか

		MaxAge: 24 * time.Hour, // preflightリクエストの結果をキャッシュする時間
	}))
	models.ConnectDatabase() // DB初期化
	models.CreateCache()     // キャッシュ初期化

	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	api := r.Group("/api")
	api.POST("/mail", controllers.SendUserEmail)
	api.POST("/login", controllers.LoginHandler)
	api.GET("/schools", controllers.FetchSchoolList)
	api.POST("/user", controllers.CreateUser)

	// 認証が必要なAPI
	api.Use(controllers.AuthMiddleware)
	{
		api.GET("/user/:id", controllers.FetchUserInfo)
		api.PUT("/user/:id", controllers.UpdateUser)
		api.GET("/family/:id", controllers.FetchFamilyById)
		api.POST("/family", controllers.CreateFamily)
		api.PUT("/family/:id", controllers.UpdateFamily)
		api.GET("/item_list/:id", controllers.FetchItemListByFamilyId)
		api.GET("/item/:id", controllers.FetchItemById)
		api.POST("/item", controllers.CreateItem)
		api.PUT("/item/:id", controllers.UpdateItem)
		api.GET("/kid_list/:id", controllers.FetchKidListByFamilyId)
		api.GET("/kid/:id", controllers.FetchKidById)
		api.POST("/kid", controllers.CreateKid)
		api.PUT("/kid/:id", controllers.UpdateKid)
		api.GET("/task_list/:id", controllers.FetchTaskListByFamilyId)
		api.GET("/task/:id", controllers.FetchTaskById)
		api.POST("/task", controllers.CreateTask)
		api.PUT("/task/:id", controllers.UpdateTask)
		api.GET("/task_type/:id", controllers.FetchTaskTypeById)
		api.GET("/task_type_list/:id", controllers.FetchTaskTypeListByFamilyId)
		api.POST("/task_type", controllers.CreateTaskType)
		api.PUT("/task_type/:id", controllers.UpdateTaskType)
		api.POST("/task_done/:id", controllers.SetTaskDone)
	}
	r.Run()
}
