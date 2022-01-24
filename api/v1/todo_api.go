package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo/common/response"
	"todo/model"
	"todo/service"
)

type TodoApi struct{}

func (t *TodoApi) CreateTodo(c *gin.Context) {
	todo := model.TodoModel{}
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		response.Fail(c, "Save new item failed: "+err.Error(), nil)
		return
	}
	var id uint
	id, err = service.TodoServiceApp.AddTodoItem(&todo)
	if err != nil {
		response.Fail(c, "Save new item failed: "+err.Error(), nil)
		return
	}
	response.Success(c, "Todo item created successfully!", gin.H{"resourceId": id})
}

func (t *TodoApi) FetchAllTodos(c *gin.Context) {
	todos, err := service.TodoServiceApp.FetchAllTodoItems()
	if err != nil {
		response.Fail(c, "Fetch all items failed :"+err.Error(), nil)
		return
	}
	if len(todos) < 1 {
		response.NotFound(c, "No todo found!", nil)
		return
	} else {
		response.Success(c, "Todos fetched!", todos)
	}
}

func (t *TodoApi) FetchSingleTodo(c *gin.Context) {
	id := c.Param("id")
	item, err := service.TodoServiceApp.FetchTodoItem(id)
	if err != nil {
		response.Fail(c, fmt.Sprintf("Fetch item %v faild: %v", id, err.Error()), nil)
		return
	}
	// Not found
	if item.ID == 0 {
		response.NotFound(c, "Item with id "+id+" not found!", nil)
		return
	}
	response.Success(c, "Todo found!", item)
}

//func (t *TodoApi) UpdateTodo(c *gin.Context) {
//	id := c.Param("id")
//
//}
