package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-api/handler"
)

func InitializeRoutes(r *gin.Engine){	
	r.GET("/opening", )
	r.GET("/opening/{id}", handler.GetOpeningById)
	r.POST("/opening", handler.CreatingOpening)
	r.PUT("/opening", )
	r.DELETE("/opening", )
}