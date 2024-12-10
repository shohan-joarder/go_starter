package controllers

import (
	"net/http"
	"strconv"

	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/services"
	"github.com/shohan-joarder/go_pos/internal/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service *services.UserService
}

// NewValidator instance for reusable validation logic
var validate = validator.New()

// NewUserController initializes a new UserController
func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

// GetAllUsers handles fetching all users
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser handles creating a new user
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	// Bind JSON input to the user model
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create user via service layer
	if err := c.service.CreateUser(&user); err != nil {
		// Check if it's a validation error
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":  "Validation failed",
				"errors": utils.FormatValidationErrors(validationErrors),
			})
			return
		}

		// Handle database or other errors
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": user})
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user models.User

	// Bind JSON input to user model
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Validate user input
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": utils.FormatValidationErrors(err)})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	user.ID = uint(id)

	// Update user in service layer
	if err := c.service.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": user})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user", "details": err.Error()})
		return
	}

	// Delete user in service layer
	if err := c.service.DeleteUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
