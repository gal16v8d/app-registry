package main

import (
	"database/sql"
	"log"

	"github.com/gal16v8d/app-registry.git/cmd/server/handler"
	"github.com/gal16v8d/app-registry.git/internal/repo"
	"github.com/gal16v8d/app-registry.git/pkg/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// @title           Go sample project
// @version         1.0
// @description     Basic crud struct.
// @termsOfService  http://swagger.io/terms/

// @contact.name   gal16v8d
// @host      localhost:8080
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	bd, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db_name")
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewSqlStore(bd)
	repoRepo := repo.NewRepository(storage)
	repoService := repo.NewService(repoRepo)
	repoHandler := handler.NewRepoHandler(repoService)

	r := gin.New()

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "Up"}) })

	repo := r.Group("/repos")
	{
		repo.GET(":id", repoHandler.GetByID())
		repo.POST("", repoHandler.CreateRepo())
		repo.PUT(":id", repoHandler.UpdateRepo())
		repo.DELETE(":id", repoHandler.DeleteRepo())
	}

	r.Run(":8080")
}
