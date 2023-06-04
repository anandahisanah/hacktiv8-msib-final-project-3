package handler

import (
	"hacktiv8-msib-final-project-3/database"
	"hacktiv8-msib-final-project-3/handler/httphandler"
	"hacktiv8-msib-final-project-3/repository/categoryrepository/categorypg"
	"hacktiv8-msib-final-project-3/repository/taskrepository/taskpg"
	"hacktiv8-msib-final-project-3/repository/userrepository/userpg"
	"hacktiv8-msib-final-project-3/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartApp() {
	db := database.GetPostgresInstance()

	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()

	userRepo := userpg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := httphandler.NewUserHandler(userService)

	categoryRepo := categorypg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := httphandler.NewCategoryHandler(categoryService)

	taskRepo := taskpg.NewTaskPG(db)
	taskService := service.NewTaskService(taskRepo, categoryRepo, userRepo)
	taskHandler := httphandler.NewTaskHandler(taskService)

	authService := service.NewAuthService(userRepo, taskRepo)

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PUT("/users/update-account", authService.Authentication(), userHandler.UpdateUser)
	r.DELETE("/users/delete-account", authService.Authentication(), userHandler.DeleteUser)

	r.POST("/categories", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.CreateCategory)
	r.GET("/categories", authService.Authentication(), categoryHandler.GetAllCategories)
	r.PUT("/categories/:categoryID", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.UpdateCategory)
	r.DELETE("/categories/:categoryID", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.DeleteCategory)

	r.POST("/tasks", authService.Authentication(), taskHandler.CreateTask)
	r.GET("/tasks", authService.Authentication(), taskHandler.GetAllTasks)
	r.PUT("/tasks/:taskID", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTask)
	r.PATCH("/tasks/update-status/:taskID", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTaskStatus)
	r.PATCH("/tasks/update-category/:taskID", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateTaskCategory)
	r.DELETE("/tasks/:taskID", authService.Authentication(), authService.TaskAuthorization(), taskHandler.DeleteTask)

	log.Fatalln(r.Run(":" + PORT))
}
