package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"tiki/dto"
	"tiki/services"
)

var path string = "password.txt"

var users []dto.UserStored

func init() {
	_, err := createFile(path)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

/*GetUser : get a UserRequest type from a username */
func GetUser(username string) (bool, dto.UserStored, error) {
	//get a user from all users
	users := GetAllUsers()
	bl, user := getStoriedUser(username, users)
	if bl == false {
		return false, user, nil
	}
	return true, user, nil
}

/*GetAllUsers : get all users UserStored in system*/
func GetAllUsers() []dto.UserStored {
	file, _ := ioutil.ReadFile(path)
	_ = json.Unmarshal([]byte(file), &users)
	return users
}

/*AddUser : add new user to system */
func AddUser(user dto.UserFromRequest) (bool, string, error) {
	//validate password
	validate, message, err := services.ValidatePassword(user.Password)
	if validate == false {
		return validate, message, err
	}
	// convert object
	newUser, err := convertToUserStoried(user)
	if err != nil {
		return false, "", err
	}
	//save
	users := GetAllUsers()
	users = append(users, newUser)
	save, err := saveFile(path, &users)
	return save, "", err
}

/*function to check a create password.txt*/
func createFile(path string) (rs bool, err error) {
	_, err = os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		defer file.Close()
		return true, err
	}
	return true, err
}

/*function to override a content to password.txt*/
func saveFile(path string, users *[]dto.UserStored) (rs bool, err error) {
	file, err := json.MarshalIndent(&users, "", " ")
	if err != nil {
		return false, err
	}
	_ = ioutil.WriteFile(path, file, 0644)
	return true, nil
}

/*function to convert UserFromRequest  -> UserStoried*/
func convertToUserStoried(user dto.UserFromRequest) (dto.UserStored, error) {
	var newUser = dto.UserStored{Username: user.Username, Password: nil}
	pass, err := services.SetNewPassword(user.Password)
	if err != nil {
		return newUser, err
	}
	newUser.Password = pass
	return newUser, nil
}

/*function find a UserStoried by username*/
func getStoriedUser(username string, users []dto.UserStored) (bool, dto.UserStored) {
	for i := len(users) - 1; i >= 0; i-- {
		if users[i].Username == username {
			return true, users[i]
		}
	}
	user := dto.UserStored{}
	return false, user
}
