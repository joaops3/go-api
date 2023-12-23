package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joaops3/go-api/pkg/dto"
	"github.com/joaops3/go-api/pkg/schemas"
)


func CreatingOpening(ctx *gin.Context){
    requestBody  := dto.CreatingOpeningDto{}

	logger.Infof("received %v",requestBody)
	

	ctx.BindJSON(&requestBody)

	if err := requestBody.Validate(); err != nil {
		logger.Errorf("validation error %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// if err := validator.New().Struct(&requestBody); err != nil {
	// 	sendError(ctx, http.StatusUnprocessableEntity, err.Error())
	// 	return
	// }
	
	opening := schemas.Opening{
		Role: requestBody.Role,
		Company: requestBody.Company,
		Location: requestBody.Location,
		Remote: *requestBody.Remote,
		Link: requestBody.Link,
		Salary: requestBody.Salary,
	}

	if err := Db.Create(&opening).Error; err != nil {
		logger.Errorf("Erro creating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	
	sendSuccess(ctx, "create", opening)
}

func GetOpeningById(ctx *gin.Context){
	id := ctx.Param("id")
	
	if id == "" {
		sendError(ctx, http.StatusBadRequest,"o id é obrigatório" )
		return
	}	
	opening := schemas.Opening{}
	if err := Db.Where("id = ?", id).Find(&opening).Error; err != nil {
		sendError(ctx, 404, "opening não encontrado" )
		return 
	}
	
	sendSuccess(ctx, "get by id", opening)
}

func GetOpeningAll(ctx *gin.Context){
	openings := []schemas.Opening{}

	if err := Db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError,"Erro ao listar" )
		return
	}
	
	sendSuccess(ctx, "get  opening", openings)
}

func UpdateOpening(ctx *gin.Context){
	newData := dto.UpdateOpeningDto{}
	fmt.Printf("received %v", newData)
	if err := ctx.BindJSON(&newData); err != nil {
		sendError(ctx, 500, "error no parse do body" )
		return 
	}
	if err := newData.Validate(); err != nil {
		sendError(ctx, 422, err.Error() )
		return 
	}
	opening := schemas.Opening{}

	Id := ctx.Param("id")
	
	if err := Db.First(&opening, Id).Error; err != nil {
		sendError(ctx, 404, "opening não encontrado" )
		return 
	}
	fmt.Printf("founded %v", opening)
	if newData.Role != "" {
		opening.Role = newData.Role
	}
	if newData.Company != "" {
		opening.Company = newData.Company
	}
	if newData.Location != "" {
		opening.Location = newData.Location
	}
	if newData.Remote != nil {
		opening.Remote = *newData.Remote
	}
	if newData.Link != "" {
		opening.Link = newData.Link
	}
	if newData.Salary > 0 {
		opening.Salary = newData.Salary
	}
	if err := Db.Save(&opening).Error; err != nil {
		logger.Errorf("erro ao salvar")
		sendError(ctx, 500, "error ao salvar" )
		return
	}
	sendSuccess(ctx, "update", opening)
}


func DeleteOpening(ctx *gin.Context){
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest,"o id é obrigatório" )
		return
	}
	opening := schemas.Opening{}
	if err := Db.First(&opening, id).Error; err != nil{
		sendError(ctx, 404, "opening não encontrado" )
		return
	}
	if err := Db.Delete(&opening).Error; err != nil {
		sendError(ctx, 500, "erro ao deletar com id" )
		return
	}
	sendSuccess(ctx, "delete", opening)
}