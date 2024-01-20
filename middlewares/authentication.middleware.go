package middlewares

import (
	// "fmt"
	"fmt"
	"net/http"
	"strings"

	"github.com/automa8e_clone/helpers"
	"github.com/gin-gonic/gin"
)

func TokenAuthenticationMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization");
	
	validTokenFormat := strings.Contains(token, "Bearer ")

	if validTokenFormat {
		token = strings.Split(token, " ")[1];
		fmt.Println("token ->", token)
		
		claims, err := helpers.ParseJWT(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Expired Access Token"})
			c.Abort()
			return
		}

		fmt.Println(claims)

		
	}

}