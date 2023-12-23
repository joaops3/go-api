package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-api/pkg/controller"
	"github.com/joaops3/go-api/pkg/middlewares"
)






func InitializeRoutes(r *gin.Engine){	
	controller.InitController()
	initializeOpeningRoutes(r)
	initializeUserRoutes(r)
	initializeAuthRoutes(r)
}

func initializeOpeningRoutes(r *gin.Engine){
	r.GET("/opening", middlewares.RequireAuth, controller.GetOpeningAll )
	r.GET("/opening/:id", middlewares.RequireAuth, controller.GetOpeningById)
	r.POST("/opening", middlewares.RequireAuth, controller.CreatingOpening)
	r.PUT("/opening/:id", middlewares.RequireAuth, controller.UpdateOpening)
	r.DELETE("/opening/:id", controller.DeleteOpening)
}

func initializeUserRoutes(r *gin.Engine){
	r.GET("/user", middlewares.RequireAuth, controller.GetUserAll )
	r.GET("/user/:id", middlewares.RequireAuth, controller.GetUserById)
	r.POST("/user",middlewares.RequireAuth, controller.CreatingUser)
	r.PUT("/user/:id", middlewares.RequireAuth, controller.UpdateUser)
	r.DELETE("/user/:id", middlewares.RequireAuth, controller.DeleteUser)
}

func initializeAuthRoutes(r *gin.Engine){
	
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	
}