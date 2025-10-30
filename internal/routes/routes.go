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
		authRoutes.GET("/test", controller.Test)

		// authRoutes.GET("/test-login",func ()  {

		// })
	}
}

// RegisterUserRoutes sets up routes for user-related operations
func RegisterUserRoutes(router *gin.RouterGroup, controller *controllers.UserController, rolePermission *services.RolePermissionService) {
	// Initialize middleware
	authMiddleware := middlewares.NewAuthorizationMiddleware(rolePermission)

	// User routes group
	userRoutes := router.Group("/users")

	userRoutes.POST("/", controller.CreateUser)

	userRoutes.Use(middlewares.AuthenticateMiddleware())
	userRoutes.Use(authMiddleware.Handle()) // Apply authorization middleware

	// Define user routes
	userRoutes.GET("/", controller.GetAllUsers)
	userRoutes.GET("/:id", controller.GetUserByID)
	// userRoutes.PUT("/:id", controller.UpdateUser)
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

func RegisterWarehouseRoutes(apiGroup *gin.RouterGroup, controller *controllers.WarehouseController, rolePermission *services.RolePermissionService) {
	// authMiddleware := middlewares.NewAuthorizationMiddleware(rolePermission)
	// Initialize middleware
	wareHouseRoutes := apiGroup.Group("warehouse")

	wareHouseRoutes.Use(middlewares.AuthenticateMiddleware())
	// wareHouseRoutes.Use(authMiddleware.Handle()) // Apply authorization middleware

	wareHouseRoutes.GET("/", controller.GetAllWarehouses)
	wareHouseRoutes.GET("/:id", controller.GetWarehouseByID)
	wareHouseRoutes.POST("/", controller.CreateWarehouse)
	wareHouseRoutes.PUT("/:id", controller.UpdateWarehouse)
	wareHouseRoutes.DELETE("/:id", controller.DeleteWarehouse)
	// authMiddleware := middlewares.NewAuthorizationMiddleware(rolePermission)

	// Product routes group
	// productRoutes := apiGroup.Group("/products",
}
