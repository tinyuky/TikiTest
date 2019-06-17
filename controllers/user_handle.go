package controllers

import (
	"tiki/dto"
	"tiki/models"
	"tiki/services"
	"github.com/gin-gonic/gin"
)

/*HandleNewUserRequest : handle request from POST /add */
func HandleNewUserRequest(ctx *gin.Context) { 
	user, err := bindDataFromRequest(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(500, err)
		return
	}
	_, msg, err := models.AddUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(500, err)
		return
	}
	if msg != "" {
		ctx.AbortWithStatusJSON(500, msg)
		return
	}
	ctx.AbortWithStatusJSON(200, "Add New User done")
}

/*HandleLoginRequest : handle request POT /login */
func HandleLoginRequest(ctx *gin.Context) {
	user, err := bindDataFromRequest(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(500, err)
		return
	}
	//get user
	bl, storedUser, err := models.GetUser(user.Username)
	if bl == false {
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.AbortWithStatusJSON(500, "Username or password not correct")
		return
	}
	//compare password with stored password
	compare, err := services.VerifyPassword(user.Password, storedUser.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(500, "Username or password not correct")
		return
	}
	if compare == false{
		ctx.AbortWithStatusJSON(500, "Username or password not correct")
		return
	} 
	ctx.AbortWithStatusJSON(200, "Login Success")
	return
}

/*HandleChangePasswordRequest handle request POST /changepassword */
func HandleChangePasswordRequest(ctx *gin.Context) {
	user, err := bindDataFromRequest(ctx)
	//Get user by username
	bl, _, err := models.GetUser(user.Username)
	if bl == false {
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.AbortWithStatusJSON(500, "Can not find user")
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(500, err)
		return
	}

	//update user
	_, msg, err := models.AddUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(500, err)
		return
	}
	if msg != "" {
		ctx.AbortWithStatusJSON(500, msg)
		return
	}
	ctx.AbortWithStatusJSON(200, "Update Password done")
}


/* function bind data from request to UserRequest type*/
func bindDataFromRequest(ctx *gin.Context) (dto.UserFromRequest, error) {
	var user dto.UserFromRequest
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}