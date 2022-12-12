//Routes/Routes.go
package Routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mooce_api/Controllers"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	grp1 := router.Group("/api")
	{
		grp1.GET("customer", Controllers.GetCustomers)
		grp1.POST("customer", Controllers.CreateCustomer)
		grp1.GET("customer/:id", Controllers.GetCustomerByID)
		grp1.PUT("customer/:id", Controllers.UpdateCustomer)
		grp1.DELETE("customer/:id", Controllers.DeleteCustomer)
	}

	return router
}
