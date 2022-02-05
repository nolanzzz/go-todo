package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"todo/global"
	"todo/model"
	"todo/model/response"
	"todo/service"
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
	if err := service.TodoServiceApp.Create(todo); err != nil {
		global.LOG.Error("todo create failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
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
	if err := service.TodoServiceApp.Update(todo); err != nil {
		global.LOG.Error("todo update failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
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
	todos, err := service.TodoServiceApp.GetAll()
	if err != nil {
		global.LOG.Error("todo get all failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
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
	items, err := service.TodoServiceApp.GetUserAll(userID)
	if err != nil {
		global.LOG.Error("todo get user all failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
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
	item, err := service.TodoServiceApp.Get(id)
	if err != nil {
		global.LOG.Error("todo get failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	// Not found
	if item.ID == 0 {
		response.NotFound(c)
		return
	}
	response.OkWithData(c, gin.H{"item": item})
}

// Done
// @Tags Todo
// @Summary Mark a task as completed
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"todo completed"}"
// @Router /api/v1/todo/done/:id [put]
func (t *TodoController) Done(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")
	if err := service.TodoServiceApp.UpdateStatus(id, userID, 1); err != nil {
		global.LOG.Error("mark todo as completed failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
	} else {
		response.OkWithMessage(c, "todo completed")
	}
}

// Undone
// @Tags Todo
// @Summary Undone a todo task
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{},"msg":"todo undone"}"
// @Router /api/v1/todo/undone/:id [put]
func (t *TodoController) Undone(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")
	if err := service.TodoServiceApp.UpdateStatus(id, userID, 0); err != nil {
		global.LOG.Error("undone todo failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
	} else {
		response.OkWithMessage(c, "todo undone")
	}
}
