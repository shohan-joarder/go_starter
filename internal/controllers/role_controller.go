package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/services"
	"github.com/shohan-joarder/go_pos/internal/utils"
	"github.com/spf13/viper"
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

func getPermissionsFilePath() (string, error) {
	exePath, err := os.Executable() // Get the path to the executable
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath) // Get the directory containing the executable
	return filepath.Join(exeDir, "data", "permissions.json"), nil
}

func (c *RoleController) PermissionsKeys(ctx *gin.Context) {

	viper.SetConfigName("permissions") // File name without extension
	viper.SetConfigType("json")        // File type
	viper.AddConfigPath("./data")      // Path to look for the configuration file
	viper.AddConfigPath(".")           // Also check the current directory

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
		return
	}

	fmt.Println("permissions", viper.Get("permissions"))

	permissions := viper.Get("permissions").(map[string]interface{})

	ctx.JSON(http.StatusOK, gin.H{"data": permissions})

	// filePath, err := getPermissionsFilePath()
	// if err != nil {
	// 	fmt.Println("Error determining file path:", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to determine file path"})
	// 	return
	// }

	// fileContent, err := os.ReadFile(filePath)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read permissions file"})
	// 	return
	// }

	// var permissions models.Permissions
	// err = json.Unmarshal(fileContent, &permissions)
	// if err != nil {
	// 	fmt.Println("Error parsing JSON:", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid permissions JSON"})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{"data": permissions})
}
