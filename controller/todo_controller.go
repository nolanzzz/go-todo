package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo/common/response"
	"todo/model"
	"todo/service/todo_service"
)

type TodoController struct{}

var Todo *TodoController

// Create
// @Tags Todo
// @Summary Create new todo task
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"todo create succeed"}"
// @Router /api/v1/todo [post]
func (t *TodoController) Create(c *gin.Context) {
	var todo model.Todo
	_ = c.ShouldBind(&todo)
	todo.UserID = c.GetUint("user_id")
	if err := todo_service.TodoServiceApp.Create(todo); err != nil {
		log.Println(err.Error())
		response.FailWithMessage(c, "todo create failed")
	} else {
		response.OkWithMessage(c, "todo create succeed")
	}
}

// Update
// @Tags Todo
// @Summary Update an existing todo task
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"todo update succeed"}"
// @Router /api/v1/todo [put]
func (t *TodoController) Update(c *gin.Context) {
	var todo model.Todo
	_ = c.ShouldBind(&todo)
	if err := todo_service.TodoServiceApp.Update(todo); err != nil {
		log.Println(err.Error())
		response.FailWithMessage(c, "todo update failed")
	} else {
		response.OkWithMessage(c, "todo update succeed")
	}
}

// GetAll
// @Tags Todo
// @Summary Get all todo tasks
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"items":{}},"msg":"succeed"}"
// @Router /api/v1/todo [get]
func (t *TodoController) GetAll(c *gin.Context) {
	todos, err := todo_service.TodoServiceApp.GetAll()
	if err != nil {
		log.Println(err.Error())
		response.FailWithMessage(c, "todo get all failed")
		return
	}
	response.OkWithData(c, gin.H{"items": todos})
}

// GetUserAll
// @Tags Todo
// @Summary Get all todo tasks of a specific user
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"items":{}},"msg":"succeed"}"
// @Router /api/v1/todo/by/:userID [get]
func (t *TodoController) GetUserAll(c *gin.Context) {
	userID := c.Param("userID")
	items, err := todo_service.TodoServiceApp.GetUserAll(userID)
	if err != nil {
		log.Println(err.Error())
		response.FailWithMessage(c, "todo get user all failed")
		return
	}
	response.OkWithData(c, gin.H{"items": items})
}

// Get
// @Tags Todo
// @Summary Get one todo task
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"item":{}},"msg":"succeed"}"
// @Router /api/v1/todo/:id [get]
func (t *TodoController) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := todo_service.TodoServiceApp.Get(id)
	if err != nil {
		log.Println(err.Error())
		response.FailWithMessage(c, "todo get failed")
		return
	}
	// Not found
	if item.ID == 0 {
		response.NotFound(c)
		return
	}
	response.OkWithData(c, gin.H{"item": item})
}
