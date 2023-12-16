package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreatingOpening(ctx *gin.Context){
    body  := struct{}{}

	ctx.BindJSON(body)
	ctx.JSON(http.StatusCreated, gin.H{"msg": "oi", "data": body})
}

func GetOpeningById(ctx *gin.Context){
	param := ctx.Params
	id := param["id"]
	ctx.JSON(http.StatusCreated, gin.H{"msg": "oi"})
}