package handler

import (
	// "hacktiv8-msib-final-project-3/database"
	// "hacktiv8-msib-final-project-3/handler/http_handler"
	// "hacktiv8-msib-final-project-3/repository/photo_repository/photo_pg"
	// "hacktiv8-msib-final-project-3/repository/user_repository/user_pg"
	// "hacktiv8-msib-final-project-3/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartApp() {
	// _ := database.GetPostgresInstance()

	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()

	log.Fatalln(r.Run(":" + PORT))
}
