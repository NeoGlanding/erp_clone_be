package middlewares

import (
	"fmt"
	"net/http"

	"github.com/automa8e_clone/helpers"

	// "fmt"

	userpartypermissions "github.com/automa8e_clone/repositories/user-party-permissions"
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