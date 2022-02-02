package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todo/common/response"
	"todo/service/todo_service"
)

type TodoController struct{}

var Todo *TodoController

// @Tags Todo
// @Summary Create new todo task
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"Successfully added new TODO item."}"
// @Router /api/v1/todo [post]
func (t *TodoController) Create(context *gin.Context) {
	title := context.PostForm("title")
	description := context.PostForm("description")
	if err := todo_service.TodoServiceApp.Create(title, description); err != nil {
		response.Fail(context, "Add new item failed: "+err.Error(), nil)
	} else {
		response.Success(context, "Successfully added new TODO item.", nil)
	}
}

// @Tags Todo
// @Summary Update an existing todo task
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"Update item successful."}"
// @Router /api/v1/todo/:id [put]
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

// @Tags Todo
// @Summary Get all todo tasks
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"items":{}},"msg":"Todos fetched!"}"
// @Router /api/v1/todo [get]
func (t *TodoController) GetAll(context *gin.Context) {
	todos, err := todo_service.TodoServiceApp.GetAll()
	if err != nil {
		response.Fail(context, "Fetch all items failed :"+err.Error(), nil)
		return
	}
	response.Success(context, "Todos fetched!", gin.H{"items": todos})
}

// @Tags Todo
// @Summary Get all todo tasks of a specific user
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"items":{}},"msg":"Fetch user's items succeed"}"
// @Router /api/v1/todo/by/:userID [get]
func (t *TodoController) GetUserAll(context *gin.Context) {
	userID := context.Param("userID")
	items, err := todo_service.TodoServiceApp.GetUserAll(userID)
	if err != nil {
		response.Fail(context, "Fetch user's items failed: "+err.Error(), nil)
		return
	}
	response.Success(context, "Fetch user's items succeed", gin.H{"items": items})
}

// @Tags Todo
// @Summary Get one todo task
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"item":{}},"msg":"Todo found!"}"
// @Router /api/v1/todo/:id [get]
func (t *TodoController) Get(context *gin.Context) {
	id := context.Param("id")
	item, err := todo_service.TodoServiceApp.Get(id)
	if err != nil {
		response.Fail(context, "Fetch item "+id+" failed: "+err.Error(), nil)
		return
	}
	// Not found
	if item.ID == 0 {
		response.NotFound(context, "Item with id "+id+" not found!", nil)
		return
	}
	response.Success(context, "Todo found!", gin.H{"item": item})
}
