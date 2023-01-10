//Routes/Routes.go
package Routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mooce_api/Controllers"
	"mooce_api/Middlewares"
	"mooce_api/Service"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	var jwtService Service.JWTService = Service.JWTAuthService()
	var loginController Controllers.LoginController = Controllers.LoginHandler(jwtService)

	grp1 := router.Group("/customer")
	{
		grp1.GET("", Middlewares.AuthorizationJWT(), Controllers.GetCustomers)
		grp1.POST("", Middlewares.AuthorizationJWT(), Controllers.CreateCustomer)
		grp1.GET("/:id", Middlewares.AuthorizationJWT(), Controllers.GetCustomerByID)
		grp1.PUT("/:id", Middlewares.AuthorizationJWT(), Controllers.UpdateCustomer)
		grp1.DELETE("/:id", Middlewares.AuthorizationJWT(), Controllers.DeleteCustomer)
	}
	grp2 := router.Group("/user")
	{
		grp2.GET("", Middlewares.AuthorizationJWT(), Controllers.GetUsers)
		grp2.POST("", Controllers.CreateUser)
		grp2.GET("/verify", loginController.VerifyUserStatus)
		grp2.GET("/:account", Middlewares.AuthorizationJWT(), Controllers.GetUserByAccount)
		grp2.PUT("/:account", Middlewares.AuthorizationJWT(), Controllers.UpdateUser)
		grp2.DELETE("/:account", Middlewares.AuthorizationJWT(), Controllers.DeleteUser)
	}

	return router
}
