package routes

import (
	"github.com/shohan-joarder/go_pos/internal/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes sets up routes for user-related operations
func RegisterUserRoutes(router *gin.RouterGroup, controller *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controller.GetAllUsers)
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.GET("/:id", controller.GetUserByID)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
	}
}
func RegisterRoleRoutes(apiGroup *gin.RouterGroup, controller *controllers.RoleController) {
	roleRoutes := apiGroup.Group("/roles")
	{
		roleRoutes.GET("/", controller.GetAllRoles)
		roleRoutes.POST("/", controller.CreateRole)
		roleRoutes.GET("/:id", controller.GetRoleByID)
		roleRoutes.PUT("/:id", controller.UpdateRole)
		roleRoutes.DELETE("/:id", controller.DeleteRole)
		// Add routes for Update, Delete...
	}
}
