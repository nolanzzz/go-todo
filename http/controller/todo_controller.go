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

func (t *TodoController) Store(context *gin.Context) {
	title := context.PostForm("title")
	description := context.PostForm("description")
	if err := todo_service.TodoServiceApp.Store(title, description); err != nil {
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

func (t *TodoController) All(context *gin.Context) {
	todos, err := todo_service.TodoServiceApp.All()
	if err != nil {
		response.Fail(context, "Fetch all items failed :"+err.Error(), nil)
		return
	}
	if len(todos) < 1 {
		response.NotFound(context, "No todo_service found!", nil)
		return
	} else {
		response.Success(context, "Todos fetched!", todos)
	}
}

func (t *TodoController) Show(context *gin.Context) {
	id := context.Param("id")
	item, err := todo_service.TodoServiceApp.Show(id)
	if err != nil {
		response.Fail(context, fmt.Sprintf("Fetch item %v faild: %v", id, err.Error()), nil)
		return
	}
	// Not found
	if item.ID == 0 {
		response.NotFound(context, "Item with id "+id+" not found!", nil)
		return
	}
	response.Success(context, "Todo found!", item)
}
