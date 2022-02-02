package todo_service

import (
	"errors"
	"todo/global"
	"todo/model"
)

type TodoService struct{}

var TodoServiceApp = new(TodoService)

func (s *TodoService) Create(title string, description string) error {
	item := model.Todo{
		Title:       title,
		Description: description,
	}
	res := global.DB.Save(&item)
	return res.Error
}

func (s *TodoService) Update(id int, title string, description string) error {
	res := global.DB.Model(&model.Todo{}).
		Where("id = ?", id).
		Update(model.Todo{Title: title, Description: description})
	if res.RowsAffected < 1 {
		return errors.New("record not found")
	}
	return res.Error
}

func (s *TodoService) GetAll() ([]model.TodoResponse, error) {
	var items []model.Todo
	res := global.DB.Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	if len(items) == 0 {
		return []model.TodoResponse{}, nil
	}
	return s.todoResponses(items), nil
}

func (s *TodoService) GetUserAll(uid int) ([]model.TodoResponse, error) {
	var items []model.Todo
	err := global.DB.Find(&items, "user_id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return s.todoResponses(items), nil
}

func (s *TodoService) todoResponses(items []model.Todo) []model.TodoResponse {
	var responses []model.TodoResponse
	// 对todo的属性做一些转换以构建更好的响应体
	for _, item := range items {
		responses = append(responses, model.TodoResponse{ID: item.ID, Title: item.Title, Completed: item.Completed})
	}
	return responses
}

func (s *TodoService) Get(id string) (model.TodoResponse, error) {
	var item model.Todo
	var resp model.TodoResponse
	if err := global.DB.First(&item, id).Error; err != nil {
		return resp, err
	}
	resp = model.TodoResponse{ID: item.ID, Title: item.Title, Completed: item.Completed}
	return resp, nil
}
