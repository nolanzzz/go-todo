package service

import (
	"todo/global"
	"todo/model"
)

type TodoService struct{}

var TodoServiceApp = new(TodoService)

func (s *TodoService) AddTodoItem(title string, completed int) (uint, error) {
	todo := model.TodoModel{
		Title:     title,
		Completed: completed,
	}
	global.DB.Save(&todo)
	return todo.ID, global.DB.Error
}

func (s *TodoService) FetchAllTodoItems() ([]model.TodoResponseModel, error) {
	var items []model.TodoModel
	var _items []model.TodoResponseModel
	result := global.DB.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(items) == 0 {
		return []model.TodoResponseModel{}, nil
	}
	// 对todo的属性做一些转换以构建更好的响应体
	for _, item := range items {
		var completed bool
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_items = append(_items, model.TodoResponseModel{ID: item.ID, Title: item.Title, Completed: completed})
	}
	return _items, nil
}

func (s *TodoService) FetchTodoItem(id string) (model.TodoResponseModel, error) {
	var item model.TodoModel
	var resp model.TodoResponseModel
	global.DB.First(&item, id)
	if global.DB.Error != nil {
		return resp, global.DB.Error
	}
	completed := false
	if item.Completed == 1 {
		completed = true
	}
	resp = model.TodoResponseModel{ID: item.ID, Title: item.Title, Completed: completed}
	return resp, nil
}

func (s *TodoService) UpdateTodoItem(id string) error {

}
