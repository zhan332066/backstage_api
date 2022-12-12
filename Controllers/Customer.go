//Controllers/Customer.go
package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mooce_api/Models"
	"net/http"
)

//GetCustomers ... Get all customer
func GetCustomers(c *gin.Context) {
	var customer []Models.Customer
	err := Models.GetAllCustomers(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//CreateCustomer ... Create Customer
func CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Models.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//GetCustomerByID ... Get the customer by id
func GetCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer Models.Customer
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//UpdateCustomer ... Update the customer information
func UpdateCustomer(c *gin.Context) {
	var customer Models.Customer
	id := c.Params.ByName("id")
	err := Models.GetCustomerByID(&customer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, customer)
	}
	c.BindJSON(&customer)
	err = Models.UpdateCustomer(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

//DeleteCustomer ... Delete the customer
func DeleteCustomer(c *gin.Context) {
	var customer Models.Customer
	id := c.Params.ByName("id")
	err := Models.DeleteCustomer(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
