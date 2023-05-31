package service

import (
	"hacktiv8-msib-final-project-3/dto"
	"hacktiv8-msib-final-project-3/entity"
	"hacktiv8-msib-final-project-3/pkg/errs"
	"hacktiv8-msib-final-project-3/repository/categoryrepository"
	"hacktiv8-msib-final-project-3/repository/taskrepository"
	"hacktiv8-msib-final-project-3/repository/userrepository"
)

type TaskService interface {
	CreateTask(user *entity.User, payload *dto.CreateTaskRequest) (*dto.CreateTaskResponse, errs.MessageErr)
	GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     taskrepository.TaskRepository
	categoryRepo categoryrepository.CategoryRepository
	userRepo     userrepository.UserRepository
}

func NewTaskService(taskRepo taskrepository.TaskRepository, categoryRepo categoryrepository.CategoryRepository, userRepo userrepository.UserRepository) TaskService {
	return &taskService{taskRepo, categoryRepo, userRepo}
}

func (t *taskService) CreateTask(user *entity.User, payload *dto.CreateTaskRequest,
) (*dto.CreateTaskResponse, errs.MessageErr) {
	task := payload.ToEntity()

	if _, checkCategoryErr := t.categoryRepo.GetCategoryByID(task.CategoryID); checkCategoryErr != nil {
		return nil, checkCategoryErr
	}

	createdTask, err := t.taskRepo.CreateTask(user, task)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateTaskResponse{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Status:      createdTask.Status,
		Description: createdTask.Description,
		UserID:      createdTask.UserID,
		CategoryID:  createdTask.CategoryID,
		CreatedAt:   createdTask.CreatedAt,
	}

	return response, nil
}

func (t *taskService) GetAllTasks() ([]dto.GetAllTasksResponse, errs.MessageErr) {
	tasks, err := t.taskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllTasksResponse{}
	for _, task := range tasks {
		user, err := t.userRepo.GetUserByID(task.UserID)
		if err != nil {
			return nil, err
		}
		response = append(response, dto.GetAllTasksResponse{
			ID:          task.CategoryID,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserID:      task.UserID,
			CategoryID:  task.CategoryID,
			CreatedAt:   task.CreatedAt,
			User: dto.UserData{
				ID:       user.ID,
				Email:    user.Email,
				FullName: user.FullName,
			},
		})
	}

	return response, nil
}
