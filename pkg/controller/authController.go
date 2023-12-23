package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joaops3/go-api/pkg/dto"
	"github.com/joaops3/go-api/pkg/schemas"
	"golang.org/x/crypto/bcrypt"
)


func SignIn(ctx *gin.Context) {
	newUser := &dto.SignInDto{}

	user := schemas.User{}
	err := Db.Where("email = ?", newUser.Email).Find(&user).Error
	if err != nil {
		sendError(ctx, 404, "Email not found")
		return
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password))

	if err != nil {
		sendError(ctx, 404, "Email or password invalid")
		return 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	
	
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {

		sendError(ctx, 500, err.Error())
		return	
	}

	resp := struct{
		 id int
		 token string}{
		id: user.ID,
		token: tokenString,
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600 * 24, "", "", false, true)

	sendSuccess(ctx, "create", resp)
}


func SignUp(ctx *gin.Context) {
	newUser := &dto.CreateUserDto{}

	ctx.BindJSON(&newUser)
	if err := newUser.Validate(); err != nil{
		sendError(ctx, 422, err.Error())
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		sendError(ctx, 422, err.Error())
		return
	}
	user := &schemas.User{
		Email: newUser.Email,
		Password: string(hashed),
	}
	if err := Db.Create(&user).Error; err != nil {
		sendError(ctx, 500, err.Error())
		return
	}
	
	sendSuccess(ctx, "create", user)
	
}