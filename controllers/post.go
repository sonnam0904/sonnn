package controllers

import (
	"github.com/astaxie/beego"
	"sonnn/models"
	"strconv"
)

type PostController struct {
	beego.Controller
}

// @router / [get]
func (o *PostController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @router / [post]
func (o *PostController) Create() {
	var post models.Post
	o.ParseForm(&post)
	result, err := models.Create(post)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}

// @router /:id [get]
func (o *PostController) GetOne() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}

	user, err := models.GetOne(uid)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = user
	}

	o.ServeJSON()
}

// @router /:id [put]
func (o *PostController) Update() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}
	id, _ := strconv.Atoi(uid)
	post := models.Post{Id: id}

	o.ParseForm(&post)
	user, err := models.Update(post)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = user
	}

	o.ServeJSON()
}

// @router /:id [delete]
func (o *PostController) Delete() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}

	id, error := strconv.Atoi(uid)
	if error != nil {
		o.Abort("500")
	}
	response, error := models.Delete(id)
	if error == nil {
		o.Data["json"] = response
	} else {
		o.Abort("500")
	}

	o.ServeJSON()
}
