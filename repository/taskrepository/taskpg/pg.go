package taskpg

import (
	"fmt"
	"hacktiv8-msib-final-project-3/entity"
	"hacktiv8-msib-final-project-3/pkg/errs"
	"hacktiv8-msib-final-project-3/repository/taskrepository"
	"log"

	"gorm.io/gorm"
)

type taskPG struct {
	db *gorm.DB
}

func NewTaskPG(db *gorm.DB) taskrepository.TaskRepository {
	return &taskPG{db}
}

func (t *taskPG) CreateTask(user *entity.User, task *entity.Task) (*entity.Task, errs.MessageErr) {
	if err := t.db.Model(user).Association("Tasks").Append(task); err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to create new task"))
	}

	return task, nil
}

func (t *taskPG) GetAllTasks() ([]entity.Task, errs.MessageErr) {
	var tasks []entity.Task
	if err := t.db.Find(&tasks).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to geet all task")
	}

	return tasks, nil
}

func (t *taskPG) GetTaskByID(id uint) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	if err := t.db.First(&task, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewNotFound(fmt.Sprintf("Task with %d is not found", id))
	}

	return &task, nil
}
