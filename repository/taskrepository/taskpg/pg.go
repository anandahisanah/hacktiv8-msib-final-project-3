package taskpg

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"hacktiv8-msib-final-project-3/entity"
	"hacktiv8-msib-final-project-3/pkg/errs"
	"hacktiv8-msib-final-project-3/repository/taskrepository"
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

func (t *taskPG) UpdateTask(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr) {
	if err := t.db.Model(oldTask).Updates(newTask).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update task with id %d", oldTask.ID))
	}

	return oldTask, nil
}

func (t *taskPG) UpdateTaskStatus(id uint, newStatus bool) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Status = newStatus

	if err := t.db.Save(task).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to update task status")
	}

	return task, nil
}

func (t *taskPG) UpdateTaskCategory(id uint, newCategoryID uint) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.CategoryID = newCategoryID

	if err := t.db.Save(task).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to update task category")
	}

	return task, nil
}

func (t *taskPG) DeleteTask(id uint) errs.MessageErr {
	if err := t.db.Delete(&entity.Task{}, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete Task with id %d", id))
	}

	return nil
}
