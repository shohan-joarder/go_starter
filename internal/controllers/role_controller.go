package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/services"
	"github.com/shohan-joarder/go_pos/internal/utils"
)

type RoleController struct {
	service *services.RoleService
}

func NewRoleController(service *services.RoleService) *RoleController {
	return &RoleController{service: service}
}

// var validate = validator.New()

func (c *RoleController) GetAllRoles(ctx *gin.Context) {
	roles, err := c.service.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	var role models.Role

	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationErrors(err)})
		return
	}

	if err := c.service.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to create role"})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

func (c *RoleController) GetRoleByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	role, err := c.service.GetRoleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) UpdateRole(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var role models.Role

	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationErrors(err)})
		return
	}

	role.ID = uint(id)

	if err := c.service.UpdateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to update role"})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) DeleteRole(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	role, err := c.service.GetRoleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.DeleteRole(role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to delete role"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
