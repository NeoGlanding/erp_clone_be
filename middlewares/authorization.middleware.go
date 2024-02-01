package middlewares

import (
	"fmt"
	"net/http"

	"github.com/automa8e_clone/helpers"

	// "fmt"

	userpartypermissions "github.com/automa8e_clone/repositories/user-party-permissions"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PartyAuthorizationRole(authorizedRole []string) (func (c *gin.Context)) {
	return func (c *gin.Context) {
		
		partyIdCtx, _ := c.Get("party-id"); partyId := partyIdCtx.(string)
		userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)


		permission, el := userpartypermissions.RetrievePermission(partyId, user["sub"].(string))

		fmt.Println("permission ->", permission)

		if el {
			permissed := helpers.StringArrayContains(authorizedRole, permission)
			if permissed {
				c.Next()
			} else {
				c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden resources"})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden resources"})
				c.Abort()
		}


	}
}

func OnboardedAuthorization(c *gin.Context) {
	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)

	userId := user["sub"].(string)

	data, isOnboarded := users_repository.CheckIsOnboarded(userId);

	if isOnboarded {
		c.Set("user-details", data)
		c.Next()
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Please onboard yourself first!"})
		c.Abort()
	}
}

func UnonboardedAuthorization(c *gin.Context) {
	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)

	userId := user["sub"].(string)

	_, isOnboarded := users_repository.CheckIsOnboarded(userId);



	if !isOnboarded {
		c.Next()
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are already onboarded!"})
		c.Abort()
	}
}