package controllers

// import (
// 	"github.com/shohan-joarder/go_pos/internal/services"
// )

// type RolePermissionController struct {
// 	service *services.RolePermissionService
// }

// func (c *RoleController) CreateRolePermission(ctx *gin.Context) {
// 	var rolePermission models.RolePermission

// 	if err := ctx.ShouldBindJSON(&rolePermission); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := validate.Struct(rolePermission); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationErrors(err)})
// 		return
// 	}

// 	if err := c.service.CreateRolePermission(&rolePermission); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to create Role Permission"})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, rolePermission)
// }
