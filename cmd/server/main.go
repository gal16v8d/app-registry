package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gal16v8d/app-registry.git/cmd/server/handler"
	"github.com/gal16v8d/app-registry.git/internal/repo"
	"github.com/gal16v8d/app-registry.git/pkg/middleware"
	"github.com/gal16v8d/app-registry.git/pkg/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// @title           Go sample project
// @version         1.0
// @description     Basic crud struct.
// @termsOfService  http://swagger.io/terms/

// @contact.name   gal16v8d
// @host      localhost:8102
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	env := os.Getenv("APP_REGISTRY_ENV")
	if env != "prod" {
		// Load the .env file in the current directory
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}
	dbUrl := os.Getenv("APP_REGISTRY_DB_URL")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewSqlStore(db)
	repoRepo := repo.NewRepository(storage)
	repoService := repo.NewService(repoRepo)
	repoHandler := handler.NewRepoHandler(repoService)

	r := gin.New()
	r.Use(middleware.AuthMiddleware())

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "Up"}) })

	repo := r.Group("/repos")
	{
		repo.GET("", repoHandler.GetAll())
		repo.GET(":id", repoHandler.GetById())
		repo.POST("", repoHandler.CreateRepo())
		repo.PUT(":id", repoHandler.UpdateRepo())
		repo.DELETE(":id", repoHandler.DeleteRepo())
	}

	r.Run(":8102")
}
