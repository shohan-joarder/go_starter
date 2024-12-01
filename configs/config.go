package configs

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shohan-joarder/go_pos/internal/controllers"
	"github.com/shohan-joarder/go_pos/internal/middlewares"
	"github.com/shohan-joarder/go_pos/internal/repositories"
	"github.com/shohan-joarder/go_pos/internal/routes"
	"github.com/shohan-joarder/go_pos/internal/services"
	"github.com/shohan-joarder/go_pos/internal/utils"
	"github.com/shohan-joarder/go_pos/pkg/database"
)

type Config struct {
	DatabaseURL string
}

func Start() {
	// Load config and database connection
	config := LoadConfig()
	db := database.ConnectPostgres(config.DatabaseURL)

	// Initialize repositories and services for roles and users
	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(roleService)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Setup router
	router := gin.Default()

	// Create the /api group
	apiGroup := router.Group("/api", middlewares.JSONValidationMiddleware())

	// Register routes for both roles and users
	routes.RegisterRoleRoutes(apiGroup, roleController)
	routes.RegisterUserRoutes(apiGroup, userController)

	// Run the server
	log.Fatal(router.Run(":8080"))
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: utils.GetEnv("DATABASE_URL", "host=localhost user=postgres password=root dbname=go_pos port=5432 sslmode=disable"),
	}
}
