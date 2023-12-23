package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joaops3/go-api/pkg/controller"
	"github.com/joaops3/go-api/pkg/schemas"
)


func RequireAuth(c *gin.Context) {
	
	// tokenString, err := c.Cookie("Authorization")

	// if err != nil {
	// 	c.AbortWithError(403, err)
	// }

	authHeader := c.GetHeader("Authorization")

		
	if authHeader == "" {
			c.AbortWithStatus(403)
			
	}

		
	if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatus(403)
			
	}

		
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if time.Now().Unix() > claims["exp"].(int64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		id := claims["sub"]

		var user schemas.User

		controller.Db.Find(&user, id)
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
	} 

	

	c.Next()
}