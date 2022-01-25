package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo/common/response"
	"todo/model"
	"todo/service/todo"
)

type TodoController struct{}

var Todo *TodoController

func (t *TodoController) Store(context *gin.Context) {
	item := model.TodoModel{}
	err := context.ShouldBindJSON(&item)
	if err != nil {
		response.Fail(context, "Save new item failed: "+err.Error(), nil)
		return
	}
	var id uint
	id, err = todo.TodoServiceApp.Store(&item)
	if err != nil {
		response.Fail(context, "Save new item failed: "+err.Error(), nil)
		return
	}
	response.Success(context, "Todo item created successfully!", gin.H{"resourceId": id})
}

func (t *TodoController) All(context *gin.Context) {
	todos, err := todo.TodoServiceApp.All()
	if err != nil {
		response.Fail(context, "Fetch all items failed :"+err.Error(), nil)
		return
	}
	if len(todos) < 1 {
		response.NotFound(context, "No todo found!", nil)
		return
	} else {
		response.Success(context, "Todos fetched!", todos)
	}
}

func (t *TodoController) Show(context *gin.Context) {
	id := context.Param("id")
	item, err := todo.TodoServiceApp.Show(id)
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

func (t *TodoController) Update(context *gin.Context) {
	//id := c.Param("id")

	item := model.TodoModel{}
	err := context.ShouldBindJSON(&item)
	if err != nil {
		response.Fail(context, "Update todo failed: "+err.Error(), nil)
		return
	}

}
