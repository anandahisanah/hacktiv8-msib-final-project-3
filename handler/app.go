package handler

import (
	"fmt"
	"hacktiv8-msib-final-project-3/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartApp() {
	db := database.GetPostgresInstance()

	fmt.Println(db)

	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()

	log.Fatalln(r.Run(":" + PORT))
}
