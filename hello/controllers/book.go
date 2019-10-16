package controllers

// import (
// 	"encoding/json"
// 	"hello/models"

// 	"github.com/astaxie/beego"
// )

// // Operations about book
// type BookController struct {
// 	beego.Controller
// }

// // @Title Create
// // @Description create book
// // @Param	body		body 	models.Book	true		"The book content"
// // @Success 200 {string} models.Book.Id
// // @Failure 403 body is empty
// // @router / [post]
// func (o *BookController) Post() {
// 	var ob models.Book
// 	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
// 	bookid := models.AddOne(ob)
// 	o.Data["json"] = map[string]string{"bookid": bookid}
// 	o.ServeJSON()
// }

// // @Title Get
// // @Description find book by bookid
// // @Param	bookid		path 	string	true		"the bookid you want to get"
// // @Success 200 {book} models.Book
// // @Failure 403 :bookid is empty
// // @router /:bookid [get]
// func (o *BookController) Get() {
// 	bookid := o.Ctx.Input.Param(":bookid")
// 	if bookid != "" {
// 		ob, err := models.GetOne(bookid)
// 		if err != nil {
// 			o.Data["json"] = err.Error()
// 		} else {
// 			o.Data["json"] = ob
// 		}
// 	}
// 	o.ServeJSON()
// }

// // @Title GetAll
// // @Description get all books
// // @Success 200 {book} models.Book
// // @Failure 403 :bookid is empty
// // @router / [get]
// func (o *BookController) GetAll() {
// 	obs := models.GetAll()
// 	o.Data["json"] = obs
// 	o.ServeJSON()
// }

// // @Title Update
// // @Description update the book
// // @Param	bookid		path 	string	true		"The bookid you want to update"
// // @Param	body		body 	models.Book	true		"The body"
// // @Success 200 {book} models.Book
// // @Failure 403 :bookid is empty
// // @router /:bookid [put]
// func (o *BookController) Put() {
// 	bookid := o.Ctx.Input.Param(":bookid")
// 	var ob models.Book
// 	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

// 	err := models.Update(bookid, ob.Score)
// 	if err != nil {
// 		o.Data["json"] = err.Error()
// 	} else {
// 		o.Data["json"] = "update success!"
// 	}
// 	o.ServeJSON()
// }

// // @Title Delete
// // @Description delete the book
// // @Param	bookid		path 	string	true		"The bookid you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 bookid is empty
// // @router /:bookid [delete]
// func (o *BookController) Delete() {
// 	bookid := o.Ctx.Input.Param(":bookid")
// 	models.Delete(bookid)
// 	o.Data["json"] = "delete success!"
// 	o.ServeJSON()
// }
