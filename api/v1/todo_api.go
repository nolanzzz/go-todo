package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"todo/global"
	"todo/model"
	"todo/model/response"
	"todo/service"
)

type TodoApi struct{}

var Todo *TodoApi

// Create
// @Tags 	 Todo
// @Summary  Create new todo task
// @Security ApiKeyAuth
// @Accept 	 application/json
// @Produce  application/json
// @Success  200 {string} string "{"status":200,"data":{},"msg":"todo create succeed"}"
// @Router   /api/v1/todo [post]
func (t *TodoApi) Create(c *gin.Context) {
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
// @Tags     Todo
// @Summary  Update an existing todo task
// @Security ApiKeyAuth
// @Accept   application/json
// @Produce  application/json
// @Success  200 {string} string "{"status":200,"data":{},"msg":"todo update succeed"}"
// @Router   /api/v1/todo [put]
func (t *TodoApi) Update(c *gin.Context) {
	var todo model.Todo
	_ = c.ShouldBind(&todo)
	if err := service.TodoServiceApp.Update(todo); err != nil {
		global.LOG.Error("todo update failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
	} else {
		response.OkWithMessage(c, "todo update succeed")
	}
}

// GetList
// @Tags Todo
// @Summary Get list of tasks separated by pages
// @Param page query int false "page number"
// @Param pageSize query int false "page size"
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"items":{},"total":int,"page":int,"pageSize":int},"msg":"succeed"}"
// @Router /api/v1/todo [get]
func (t *TodoApi) GetList(c *gin.Context) {
	var page, pageSize int
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	total, err := service.TodoServiceApp.TotalCount(0)
	if err != nil {
		global.LOG.Error("get total todo count failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	todos, err := service.TodoServiceApp.GetList(page, pageSize)
	if err != nil {
		global.LOG.Error("get todo list failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, gin.H{
		"items":    todos,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetListByUser
// @Tags Todo
// @Summary Get list of tasks separated by pages from user
// @Param page query int false "page number"
// @Param pageSize query int false "page size"
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"items":{},"total":int,"page":int,"pageSize":int},"msg":"succeed"}"
// @Param   userID path int true "id of user"
// @Router /api/v1/todo/by/:userID [get]
func (t *TodoApi) GetListByUser(c *gin.Context) {
	var page, pageSize int
	userID, _ := strconv.Atoi(c.Param("userID"))
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	total, err := service.TodoServiceApp.TotalCount(userID)
	if err != nil {
		global.LOG.Error("get user's todo count failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	todos, err := service.TodoServiceApp.GetListByUser(userID, page, pageSize)
	if err != nil {
		global.LOG.Error("get todo list by user failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, gin.H{
		"items":    todos,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// Get
// @Tags Todo
// @Summary Get one todo task
// @Produce application/json
// @Success 200 {string} string "{"status":200,"data":{"item":{}},"msg":"succeed"}"
// @Param   id path int true "id of task"
// @Router /api/v1/todo/:id [get]
func (t *TodoApi) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.LOG.Error("invalid item id", zap.Error(err))
		response.FailWithMessage(c, "invalid item id")
		return
	}
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
// @Param   id path int true "id of task"
// @Router /api/v1/todo/done/:id [put]
func (t *TodoApi) Done(c *gin.Context) {
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
// @Param   id path int true "id of task"
// @Router /api/v1/todo/undone/:id [put]
func (t *TodoApi) Undone(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")
	if err := service.TodoServiceApp.UpdateStatus(id, userID, 0); err != nil {
		global.LOG.Error("undone todo failed", zap.Error(err))
		response.FailWithMessage(c, err.Error())
	} else {
		response.OkWithMessage(c, "todo undone")
	}
}
