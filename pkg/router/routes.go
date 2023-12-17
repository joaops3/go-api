package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-api/controller"
)

func InitializeRoutes(r *gin.Engine){	
	controller.InitController()
	r.GET("/opening", controller.GetOpeningAll )
	r.GET("/opening/:id", controller.GetOpeningById)
	r.POST("/opening", controller.CreatingOpening)
	r.PUT("/opening/:id", controller.UpdateOpening)
	r.DELETE("/opening/:id", controller.DeleteOpening)
}