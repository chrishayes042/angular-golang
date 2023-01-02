package model

import (
	"fmt"
)

type User struct{
	UserId int
	UserName string
	UserPass string
	UserEmail string

}



func GetUserList(User) []User{
	var userList []User

	userList = append(userList, User{})
	fmt.Println("The userList is adding")
	return userList
}
func (e User) GetUserInfo(id int, name string, email string){
	fmt.Printf("User ID: %d\n", id)
	fmt.Printf("User Name: %s\n", name)
	fmt.Printf("User Email: %s\n", email)
	// fmt.Println("Hi from the model package.")
}