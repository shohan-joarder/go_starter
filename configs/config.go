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

	// Initialize repositories and services for role permissions
	rolePermissionRepo := repositories.NewRolePermissionRepository(db)
	rolePermissionService := services.NewRolePermissionService(rolePermissionRepo)

	// Initialize repositories and services for roles
	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(roleService)

	// Initialize repositories and services for users
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Initialize repositories and services for auth
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	// Setup router
	router := gin.Default()

	// Create the /api group
	apiGroup := router.Group("/api", middlewares.JSONValidationMiddleware())

	// Register routes for roles, users, and auth
	routes.RegisterRoleRoutes(apiGroup, roleController, rolePermissionService) // Role routes
	routes.RegisterUserRoutes(apiGroup, userController, rolePermissionService) // User routes
	routes.RegisterAuthRoutes(apiGroup, authController)                        // Auth routes

	// Run the server
	log.Fatal(router.Run(":8080"))
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: utils.GetEnv("DATABASE_URL", "host=localhost user=postgres password=root dbname=go_pos port=5432 sslmode=disable"),
	}
}
