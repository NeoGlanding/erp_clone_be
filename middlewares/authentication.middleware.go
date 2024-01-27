package middlewares

import (
	"net/http"
	"strings"

	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/models"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/gin-gonic/gin"
)

func TokenAuthenticationMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization");
	var user models.User
	
	validTokenFormat := strings.Contains(token, "Bearer ")

	if validTokenFormat {
		token = strings.Split(token, " ")[1];
		
		claims, err := helpers.ParseJWT(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Expired Access Token"})
			c.Abort()
			return
		}

		if claims["ic"] == nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Expired Access Token"})
			c.Abort()
			return
		} else {
			users_repository.FindByEmail(claims["email"].(string), &user);

			if claims["ic"].(float64) != float64(user.InformationChanged) {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Expired Access Token"})
				c.Abort()
				return
			}
		}

		
		c.Set("user", claims)
		
	} else {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Unauthorized"})
		c.Abort()
	}

}