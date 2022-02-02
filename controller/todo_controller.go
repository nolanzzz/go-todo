package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"todo/common/response"
	"todo/service/todo_service"
)

type TodoController struct{}

var Todo *TodoController

func (t *TodoController) Create(context *gin.Context) {
	title := context.PostForm("title")
	description := context.PostForm("description")
	if err := todo_service.TodoServiceApp.Create(title, description); err != nil {
		response.Fail(context, "Add new item failed: "+err.Error(), nil)
	} else {
		response.Success(context, "Successfully added new TODO item.", nil)
	}
}

func (t *TodoController) Update(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	title := context.PostForm("title")
	description := context.PostForm("description")
	if err := todo_service.TodoServiceApp.Update(id, title, description); err != nil {
		response.Fail(context, "Update item failed: "+err.Error(), nil)
	} else {
		response.Success(context, "Update item successful.", nil)
	}
}

func (t *TodoController) GetAll(context *gin.Context) {
	todos, err := todo_service.TodoServiceApp.GetAll()
	if err != nil {
		response.Fail(context, "Fetch all items failed :"+err.Error(), nil)
		return
	}
	response.Success(context, "Todos fetched!", todos)
}

func (t *TodoController) GetUserAll(context *gin.Context) {
	uid := context.GetInt("user_id")
	items, err := todo_service.TodoServiceApp.GetUserAll(uid)
	if err != nil {
		response.Fail(context, "Fetch user's items failed: "+err.Error(), nil)
		return
	}
	response.Success(context, "Fetch user's items succeed", gin.H{"items": items})
}

func (t *TodoController) Get(context *gin.Context) {
	id := context.Param("id")
	item, err := todo_service.TodoServiceApp.Get(id)
	if err != nil {
		response.Fail(context, fmt.Sprintf("Fetch item %v faild: %v", id, err.Error()), nil)
		return
	}
	// Not found
	if item.ID == 0 {
		response.NotFound(context, "Item with id "+id+" not found!", nil)
		return
	}
	response.Success(context, "Todo found!", gin.H{"item": item})
}
