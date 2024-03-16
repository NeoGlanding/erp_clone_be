package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Customer(c *gin.Engine) {
	customer := c.Group("/customers")

	customer.GET("/", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization, middlewares.InterceptPartyIdFromQueryRequired,
		middlewares.PartyAuthorizationRole([]string{"admin", "owner", "viewer"}), middlewares.PaginationMiddleware,
		middlewares.QueryMiddleware, controllers.GetCustomers, middlewares.ResponseMiddlewares)

	customer.GET("/:id", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization, middlewares.InterceptPartyIdFromQueryRequired,
		middlewares.PartyAuthorizationRole([]string{"admin", "owner", "viewer"}), controllers.GetCustomer, middlewares.ResponseMiddlewares)

	customer.GET("/types", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization, controllers.GetCustomerType, middlewares.ResponseMiddlewares)
	customer.GET("/partnerships", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization, controllers.GetCustomerPartnership, middlewares.ResponseMiddlewares)

	customer.POST("/", middlewares.TokenAuthenticationMiddleware,
		middlewares.OnboardedAuthorization, middlewares.CustomerTypeIdExistBody, middlewares.CustomerPartnershipIdExistBody,
		middlewares.BodyCountryIdExistMiddleware, middlewares.InterceptFileIdBody("file_id"), middlewares.FileIdExist,
		middlewares.InterceptPartyIdFromBody, middlewares.PartyAuthorizationRole([]string{"admin", "owner"}),
		controllers.CreateCustomer, middlewares.ResponseMiddlewares)

	customer.PUT("/:id", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization,
		middlewares.CustomerTypeIdExistBody, middlewares.CustomerPartnershipIdExistBody, middlewares.BodyCountryIdExistMiddleware,
		middlewares.InterceptFileIdBody("file_id"), middlewares.FileIdExist, middlewares.InterceptPartyIdFromQueryRequired,
		middlewares.PartyAuthorizationRole([]string{"admin", "owner"}),
		controllers.UpdateCustomer, middlewares.ResponseMiddlewares)

	customer.POST("/:id/addresses", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization,
		middlewares.InterceptPartyIdFromQueryRequired, middlewares.PartyAuthorizationRole([]string{"admin", "owner"}),
		controllers.CreateCustomerAddress, middlewares.ResponseMiddlewares)

	customer.POST("/:id/contacts", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization,
		middlewares.InterceptPartyIdFromQueryRequired, middlewares.PartyAuthorizationRole([]string{"admin", "owner"}),
		controllers.CreateContacts, middlewares.ResponseMiddlewares)

	customer.PUT("/:id/addresses", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization,
		middlewares.InterceptPartyIdFromQueryRequired, middlewares.PartyAuthorizationRole([]string{"admin", "owner"}),
		controllers.UpdateCustomerAddresses, middlewares.ResponseMiddlewares)

	customer.PUT("/:id/contacts", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization,
		middlewares.InterceptPartyIdFromQueryRequired, middlewares.PartyAuthorizationRole([]string{"admin", "owner"}),
		controllers.UpdateContacts, middlewares.ResponseMiddlewares)

}
