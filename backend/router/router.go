package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	_ "taschola/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"taschola/controllers"
)

func Init() *gin.Engine {
	// initialize engine
	engine := gin.Default()

	// CORS settings
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health check
	engine.GET("/health", controllers.HealthCheck)
	engine.GET("/healthz", controllers.CheckDBConnection)

	// endpoints version 1
	v1 := engine.Group("v1")
	{
		// tasks
		v1.GET("/tasks", controllers.GetTasks)
		v1.GET("/t2schola_sync", controllers.NotImplemented)

		task := v1.Group("/task")
		// task.Use(controllers.LoginCheck)
		{
			task.GET("/:id", controllers.GetTask)
			task.PUT("/:id", controllers.UpdateTask)
			task.DELETE("/:id", controllers.DeleteTask)
			task.POST("/new", controllers.CreateTask)
		}

		// auth
		v1.POST("/login", controllers.Login)
		v1.POST("/logout", controllers.Logout)

		// user
		v1.POST("/user/new", controllers.CreateUser)

		user := v1.Group("/user")
		// user.Use(controllers.LoginCheck)
		{
			user.GET("/:id", controllers.GetUser)
			user.PUT("/:id", controllers.UpdateUser)    // edit
			user.DELETE("/:id", controllers.DeleteUser) // delete
		}
	}

	return engine
}
