package service

import (
	"errors"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
	"todo/global"
	"todo/model"
	"todo/utils"
)

type TodoService struct{}

var TodoServiceApp = new(TodoService)

func (s *TodoService) Create(todo model.Todo) error {
	err := global.DB.Create(&todo).Error
	return err
}

func (s *TodoService) Update(todo model.Todo) error {
	res := global.DB.Model(&todo).Where("id = ?", todo.ID).Updates(todo)
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

func (s *TodoService) GetUserAll(userID string) ([]model.TodoResponse, error) {
	var items []model.Todo
	var user model.User
	err := global.DB.Find(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	if err = global.DB.Model(&user).Association("Todos").Find(&items); err != nil {
		return nil, err
	}
	return s.todoResponses(items), nil
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

func (s *TodoService) UpdateStatus(id string, userID uint, status int) error {
	var todo model.Todo
	if err := global.DB.Find(&todo, "id = ?", id).Error; err != nil {
		return err
	}
	if todo.UserID != userID {
		return errors.New("user not authorized")
	}
	if todo.Completed == status {
		return nil
	}
	todo.Completed = status
	// calculate time spent on complete
	if status == 1 {
		todo.TimeSpent = utils.TimeDiffMinutes(todo.CreatedAt, time.Now())
	} else {
		todo.TimeSpent = 0
	}
	if err := global.DB.Save(&todo).Error; err != nil {
		return err
	}
	var user model.User
	if err := global.DB.First(&user, "id = ?", userID).Error; err != nil {
		global.LOG.Error("user not found")
	}
	// Update ranking stats in redis
	if err := s.updateRedisStats(user.Username, todo.TimeSpent); err != nil {
		global.LOG.Error("update redis stats failed", zap.Error(err))
	}
	return nil
}

func (s *TodoService) updateRedisStats(username string, timeSpent int) (err error) {
	err = global.REDIS.ZIncrBy(global.CONFIG.Redis.KeyRankMinutes, float64(timeSpent), username).Err()
	if err != nil {
		return err
	}
	err = global.REDIS.ZIncrBy(global.CONFIG.Redis.KeyRankTodos, 1, username).Err()
	return err
}

func (s *TodoService) todoResponses(items []model.Todo) []model.TodoResponse {
	var responses []model.TodoResponse
	// 对todo的属性做一些转换以构建更好的响应体
	for _, item := range items {
		responses = append(responses, model.TodoResponse{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Completed:   item.Completed,
		})
	}
	return responses
}

func (s *TodoService) GenerateTodos(count int) {
	// Create new tasks for demo users
	var userID int
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		userID = rand.Intn(9) + 1
		global.LOG.Debug("userID", zap.Int("userID", userID))
		title := "Todo title at " + time.Now().Format("2006-01-02 15:04:05")
		description := "Description at " + time.Now().Format("2006-01-02 15:04:05")
		err := s.Create(model.Todo{Title: title, Description: description, UserID: uint(userID)})
		if err != nil {
			global.LOG.Error("generating todo failed", zap.Error(err))
		}
	}
}

func (s *TodoService) CompleteTodos() {
	// Randomly pick a demo user, and complete all their todos
	var todos []model.Todo
	var err error
	rand.Seed(time.Now().Unix())
	userID := rand.Intn(9) + 1
	if err = global.DB.Find(&todos, "user_id = ?", userID).Error; err != nil {
		global.LOG.Error("CompleteTodos failed", zap.Error(err))
		return
	}
	for _, todo := range todos {
		if err = s.UpdateStatus(strconv.Itoa(int(todo.ID)), todo.UserID, 1); err != nil {
			global.LOG.Error("CompleteTodos failed", zap.Error(err))
		}
	}
}
