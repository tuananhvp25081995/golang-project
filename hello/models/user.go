package models

import (
	"fmt"
	"hello/libs"
	"log"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	UserList []orm.Params
	Data     []User
)

type User struct {
	Id         int    `orm:"column(id);pk";"AUTO_INCREMENT"`
	Username   string `orm:"column(username);unique"`
	Email      string `orm:"column(email);unique"`
	Password   string `orm:"column(password)"`
	Fullname   string `orm:"column(fullname);null"`
	Birthday   string `orm:"column(birthday);null"`
	Gender     string `orm:"column(gender);null"`
	Age        string `orm:"column(age);null"`
	Salt       string `orm:"column(salt)"`
	IsVerified int    `orm:"column(isVerified)"`
}

type Token struct {
	Id    int    `orm:"column(id);pk"`
	Email string `orm:"column(email);unique"`
	Token string `orm:"column(token);unique"`
}

func (t *Token) TableName() string {
	return "tokens"
}
func init() {
	orm.RegisterModel(new(Token))
}

func (t *User) TableName() string {
	return "users"
}
func init() {
	orm.RegisterModel(new(User))
}
func AddUser(param User) error {
	o := orm.NewOrm()
	tempVar := User{
		// Id:       param.Id,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
		Fullname: param.Fullname,
		Salt:     param.Salt,
	}
	var tokens = strings.ToUpper(libs.GetRandomString(6))
	var newToken = Token{
		// Id:       param.Id,
		Email: param.Email,
		Token: tokens,
	}

	from := beego.AppConfig.String("database")
	pass := beego.AppConfig.String("pass")
	to := param.Email
	msg := []byte("From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Account Verification Token" +
		"\r\n" + "Dear" + " " + param.Email + "," + "\n\n\n" +
		"Here is the account activation confirmation code:\n\n" + tokens)

	errs := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if errs != nil {
		log.Printf("smtp error: %s", errs)
		return errs
	}

	var user User
	err := o.QueryTable("users").Filter("username", param.Username).One(&user)
	if err == nil {
		return orm.ErrNoRows
	}
	token, errors := o.Insert(&newToken)
	id, err := o.Insert(&tempVar)
	if errors == nil {
		fmt.Println(token)
	}
	if err == nil {
		fmt.Println(id)
	}
	return nil

}

func ConfirmEmail(token Token) bool {
	o := orm.NewOrm()
	// var r orm.RawSeter
	var Token = token.Token
	var maps []orm.Params
	o.Raw("SELECT token,email FROM tokens").Values(&maps)
	if Token != maps[0]["token"] {
		return false
	}
	res, err := o.Raw("UPDATE users SET isVerified = ? WHERE email = ?", 1, maps[0]["email"]).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return true
}

func GetUser(username string) (u map[string]interface{}, err error) {
	o := orm.NewOrm()
	var user User
	var maps []orm.Params
	errs := o.QueryTable("users").Filter("username", username).One(&user)
	if errs == nil {
		o.Raw("SELECT username,email,fullname,birthday,gender,age FROM users WHERE username = ?", username).Values(&maps)
		info := map[string]interface{}{
			"username": fmt.Sprintf("%v", maps[0]["username"]),
			"email":    fmt.Sprintf("%v", maps[0]["email"]),
			"fullname": fmt.Sprintf("%v", maps[0]["fullname"]),
			"birthday": fmt.Sprintf("%v", maps[0]["birthday"]),
			"gender":   fmt.Sprintf("%v", maps[0]["gender"]),
			"age":      fmt.Sprintf("%v", maps[0]["age"]),
		}
		return info, nil
	}
	return nil, err
}

func GetAllUsers() []orm.Params {
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("SELECT username, fullname,email,gender,age FROM users").Values(&UserList)
	fmt.Println(UserList)
	return UserList
}

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	o := orm.NewOrm()
// 	user := User{Id: uid}
// 	if o.Read(&user) == nil {
// 		user.Username = uu.Username
// 		user.Password = uu.Password
// 		if num, err := o.Update(&user); err == nil {
// 			fmt.Println(num)
// 		}
// 	}
// 	return &user, nil
// }

// func DeleteUser(uid string) {
// 	o := orm.NewOrm()
// 	if num, err := o.Delete(&User{Id: uid}); err == nil {
// 		fmt.Println(num)
// 	}
// }

func Login(username, password string) bool {
	o := orm.NewOrm()
	var maps []orm.Params
	var user User
	err := o.QueryTable("users").Filter("username", username).One(&user)
	if err == orm.ErrNoRows {
		return false
	}
	o.Raw("SELECT id,isVerified FROM users WHERE username = ?", username).Values(&maps)
	id := (maps[0]["id"])
	str := fmt.Sprintf("%v", id)
	isVerified := maps[0]["isVerified"]
	str1 := fmt.Sprintf("%v", isVerified)
	i2, errs := strconv.Atoi(str1)
	if errs == nil {
		if i2 == 1 {
			i1, err := strconv.Atoi(str)
			if err == nil {
				findUser := User{Id: i1}
				o.Read(&findUser)
				if findUser.Password == libs.Md5([]byte(password+findUser.Salt)) {
					return true
				}
			}
		}
	}
	return false
}

// func Checkusername(username string) bool {
// 	o := orm.NewOrm()
// 	var user User
// 	err := o.QueryTable("users").Filter("username", username).One(&user)
// 	if err == orm.ErrNoRows {
// 		return false
// 	}
// 	return true
// }

// func Resetpasword(uu *User) bool {
// 	o := orm.NewOrm()
// 	var maps []orm.Params
// 	// fmt.Println(uu.Username)
// 	o.Raw("SELECT id FROM users WHERE username = ?", uu.Username).Values(&maps)
// 	id := maps[0]["id"]
// 	str := fmt.Sprintf("%v", id)
// 	findUser := User{Id: str}
// 	o.Read(&findUser)
// 	findUser.Password = uu.Password
// 	findUser.Salt = uu.Salt
// 	// fmt.Println(findUser.Password)
// 	if num, err := o.Update(&findUser); err == nil {
// 		fmt.Println(num)
// 	}
// 	return true
// }
