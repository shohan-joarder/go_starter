package routes

import (
	"github.com/shohan-joarder/go_pos/internal/controllers"
	"github.com/shohan-joarder/go_pos/internal/middlewares"
	"github.com/shohan-joarder/go_pos/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(apiGroup *gin.RouterGroup, controller *controllers.AuthController) {
	authRoutes := apiGroup.Group("/auth")
	{
		authRoutes.POST("/login", controller.Login)
	}
}

// RegisterUserRoutes sets up routes for user-related operations
func RegisterUserRoutes(router *gin.RouterGroup, controller *controllers.UserController, rolePermission *services.RolePermissionService) {
	// Initialize middleware
	authMiddleware := middlewares.NewAuthorizationMiddleware(rolePermission)

	// User routes group
	userRoutes := router.Group("/users",
		middlewares.AuthenticateMiddleware(), // Apply authentication middleware
		authMiddleware.Handle(),              // Apply authorization middleware
	)

	// Define user routes
	userRoutes.POST("/", controller.CreateUser)
	userRoutes.GET("/:id", controller.GetUserByID)
	userRoutes.PUT("/:id", controller.UpdateUser)
	userRoutes.GET("/", controller.GetAllUsers)
	userRoutes.DELETE("/:id", controller.DeleteUser)
}

func RegisterRoleRoutes(apiGroup *gin.RouterGroup, controller *controllers.RoleController, rolePermission *services.RolePermissionService) {
	// Initialize middleware
	authMiddleware := middlewares.NewAuthorizationMiddleware(rolePermission)

	// Role routes group
	roleRoutes := apiGroup.Group("/roles",
		middlewares.AuthenticateMiddleware(), // Apply authentication middleware
		authMiddleware.Handle(),              // Apply authorization middleware
	)

	// Define role routes
	roleRoutes.GET("/", controller.GetAllRoles)
	roleRoutes.POST("/", controller.CreateRole)
	roleRoutes.GET("/:id", controller.GetRoleByID)
	roleRoutes.PUT("/:id", controller.UpdateRole)
	roleRoutes.DELETE("/:id", controller.DeleteRole)
}
