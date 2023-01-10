//Models/UserModel.go
package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"mooce_api/Config"
)

type User struct {
	Id         uint   `json:"id"`
	Account    string `json:"account"`
	Pwd        string `json:"pwd"`
	Salt       string `json:"salt"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
	Status     uint   `json:"status"`
	Note       string `json:"note"`
}

func (b *User) TableName() string {
	return "User"
}

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	userPwd := []byte(user.Pwd)
	HashPwd, err := bcrypt.GenerateFromPassword(userPwd, 10)
	if err != nil {
		return err
	}
	user.Pwd = string(HashPwd)
	fmt.Println("length", len(string(HashPwd)))
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by account
func GetUserByAccount(user *User, account string) (err error) {
	if err = Config.DB.Where("Account = ?", account).First(user).Error; err != nil {
		fmt.Println("Account", err)
		return err
	}
	return nil
}

//VerifyUserStatus ... Fetch only one user by account
func VerifyUserStatus(user *User, account string, pwd string) (err error) {

	if err = Config.DB.Where("Account = ?", account).First(user).Error; err != nil {

		return err
	}

	byteHashpwd := []byte(user.Pwd)
	bytepwd := []byte(pwd)

	err = bcrypt.CompareHashAndPassword(byteHashpwd, bytepwd)
	if err != nil {
		fmt.Println("Error")

		return err
	}

	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, account string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, account string) (err error) {
	Config.DB.Where("account = ?", account).Delete(user)
	return nil
}
