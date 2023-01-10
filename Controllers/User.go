//Controllers/User.go
package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mooce_api/Models"
	"mooce_api/Service"
	"net/http"
)

//GetUsers ... Get all user
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByAccount ... Get the user by account
func GetUserByAccount(c *gin.Context) {
	account := c.Params.ByName("account")
	var user Models.User
	err := Models.GetUserByAccount(&user, account)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//login contorller interface
type LoginController interface {
	VerifyUserStatus(c *gin.Context)
}

type loginController struct {
	jWtService Service.JWTService
}

func LoginHandler(jWtService Service.JWTService) LoginController {
	return &loginController{
		jWtService: jWtService,
	}
}

//VerifyUserStatus ... Verify UserData
func (controller *loginController) VerifyUserStatus(c *gin.Context) {
	account := c.Query("account")
	pwd := c.Query("pwd")

	var user Models.User
	err := Models.VerifyUserStatus(&user, account, pwd)
	fmt.Println("result", err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
	} else {
		token := controller.jWtService.GenerateToken(account, false)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	account := c.Params.ByName("account")
	err := Models.GetUserByAccount(&user, account)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, account)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	account := c.Params.ByName("account")
	err := Models.DeleteUser(&user, account)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"account" + account: "is deleted"})
	}
}
