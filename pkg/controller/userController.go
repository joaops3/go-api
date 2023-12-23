package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joaops3/go-api/pkg/dto"
	"github.com/joaops3/go-api/pkg/schemas"
)


func CreatingUser(ctx *gin.Context){
	newUser := &dto.CreateUserDto{}

	ctx.BindJSON(&newUser)
	if err := newUser.Validate(); err != nil{
		sendError(ctx, 422, err.Error())
		return
	}

	user := &schemas.User{
		Email: newUser.Email,
		Password: newUser.Password,
	}
	if err := Db.Create(&user).Error; err != nil {
		sendError(ctx, 500, err.Error())
		return
	}
	
	sendSuccess(ctx, "create", user)
}

func GetUserById(ctx *gin.Context){
	id := ctx.Param("id")
	
	 user := Db.Find("id = ?", id) 
	 if user == nil {
		sendError(ctx, 404, "User not found with id")
		return
	 }
	 
	sendSuccess(ctx, "get by id", user)
}

func GetUserAll(ctx *gin.Context){
	Users := []schemas.User{}

	if err := Db.Find(&Users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,"Erro ao listar" )
		return
	}
	
	sendSuccess(ctx, "get  User", Users)
}

func UpdateUser(ctx *gin.Context){

}


func DeleteUser(ctx *gin.Context){
	 currentLoggedUser, _ := ctx.Get("user")

	 if currentLoggedUser == nil {
		sendError(ctx, 404, "current user not found")
	 }
	
}