package controllers

import (
	"encoding/json"
	"fmt"
	"hello/libs"
	"hello/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body  models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /signup [post]
func (u *UserController) Signup() {
	var param models.User
	// param.Id = u.GetString("id")
	param.Username = u.GetString("username")
	param.Email = u.GetString("email")
	param.Password = u.GetString("password")
	pwd, salt := libs.Password(4, param.Password)
	param.Password = pwd
	param.Salt = salt
	param.Fullname = u.GetString("fullname")
	var err error
	err = models.AddUser(param)
	fmt.Println(err)
	if err != nil {
		u.Data["json"] = map[string]interface{}{"status": "user already exists"}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]interface{}{"status": "A verification token has been sent to " + param.Username + "."}
	u.ServeJSON()
	return
}

// @Title ConfirmEmail
// @Description confirm email
// @Success 200 {object} models.Token
// @router /confirmemail [post]
func (u *UserController) ConfirmEmail() {
	var token models.Token
	token.Token = u.GetString("token")
	var bol bool
	bol = models.ConfirmEmail(token)
	if bol == false {
		u.Data["json"] = map[string]interface{}{"status": "We were unable to find a user for this token."}
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]interface{}{"status": "The account has been verified. Please log in."}
	u.ServeJSON()
	return
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	// userss := users[0]["username"]
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by username
// @Param	username		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :username is empty
// @router /:username [get]
func (u *UserController) Get() {
	username := u.GetString(":username")
	if username != "" {
		user, err := models.GetUser(username)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// // @Title Update
// // @Description update the user
// // @Param	uid		path 	string	true		"The uid you want to update"
// // @Param	body		body 	models.User	true		"body for user content"
// // @Success 200 {object} models.User
// // @Failure 403 :uid is not int
// // @router /:uid [put]
// func (u *UserController) Put() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		var User models.User
// 		User.Username = u.GetString("username")
// 		User.Password = u.GetString("password")
// 		pwd, salt := libs.Password(4, "")
// 		User.Password = pwd
// 		User.Salt = salt
// 		uu, err := models.UpdateUser(uid, &User)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = uu
// 		}
// 	}
// 	u.ServeJSON()
// }

// // @Title Delete
// // @Description delete the user
// // @Param	uid		path 	string	true		"The uid you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 uid is empty
// // @router /:uid [delete]
// func (u *UserController) Delete() {
// 	uid := u.GetString(":uid")
// 	models.DeleteUser(uid)
// 	u.Data["json"] = "delete success!"
// 	u.ServeJSON()
// }

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	o := orm.NewOrm()
	var maps []orm.Params
	var Users models.User
	username := u.GetString("username")
	password := u.GetString("password")
	o.Raw("SELECT * FROM users WHERE username = ?", username).Values(&maps)
	if models.Login(username, password) {
		json.Unmarshal(u.Ctx.Input.RequestBody, &Users)
		token := models.AddToken(Users, u.Ctx.Input.Domain())
		u.Data["json"] = map[string]interface{}{
			"username":  maps[0]["username"],
			"email":     maps[0]["email"],
			"birthday":  maps[0]["birthday"],
			"gender":    maps[0]["gender"],
			"age":       maps[0]["age"],
			"fullnname": maps[0]["fullname"],
			"token":     token,
		}
	} else {
		u.Data["json"] = "The username or password is incorrect"
	}
	u.ServeJSON()
}

// // @Title Resetpassword
// // @Description Resetpassword user into the system
// // @Param	username		query 	string	true		"The username for resetpassword"
// // @Param	newpassword		query 	string	true		"The newpassword for resetpassword"
// // @Success 200 {string} re success
// // @Failure 403 user not exist
// // @router /resetpassword [post]
// func (u *UserController) Resetpassword() {
// 	var Users models.User
// 	username := u.GetString("username")
// 	newpassword := u.GetString("password")
// 	if models.Checkusername(username) {
// 		Users.Password = newpassword
// 		Users.Username = username
// 		pwd, salt := libs.Password(4, "")
// 		Users.Password = pwd
// 		Users.Salt = salt
// 		if models.Resetpasword(&Users) {
// 			// fmt.Println(username, newpassword)
// 			u.Data["json"] = "Resetpassword success"
// 		}
// 	} else {
// 		u.Data["json"] = "User not exits"
// 	}
// 	u.ServeJSON()
// }

// // @Title logout
// // @Description Logs out current logged in user session
// // @Success 200 {string} logout success
// // @router /logout [get]
// func (u *UserController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }
