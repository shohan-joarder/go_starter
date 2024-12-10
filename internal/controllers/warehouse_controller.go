package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/services"
	"github.com/shohan-joarder/go_pos/internal/utils"
)

type WarehouseController struct {
	warehouseService *services.WarehouseService
}

// Constructor
func NewWarehouseController(warehouseService *services.WarehouseService) *WarehouseController {
	return &WarehouseController{warehouseService: warehouseService}
}

func (c *WarehouseController) GetAllWarehouses(ctx *gin.Context) {
	warehouses, err := c.warehouseService.GetAllWarehouses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve warehouses", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": warehouses})
}

func (c *WarehouseController) GetWarehouseByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	warehouse, err := c.warehouseService.GetWarehouseByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve warehouse", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": warehouse})
}

func (c *WarehouseController) CreateWarehouse(ctx *gin.Context) {
	var warehouse models.Warehouse

	// Bind JSON input to the warehouse struct
	if err := ctx.ShouldBindJSON(&warehouse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Extract user ID from context
	userID, err := utils.ParseUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID",
			"details": err.Error(),
		})
		return
	}

	// Assign user ID to warehouse
	warehouse.UserID = userID

	// Validate the warehouse input
	if err := validate.Struct(warehouse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": utils.FormatValidationErrors(err),
		})
		return
	}

	// Call the service to create the warehouse
	if err := c.warehouseService.CreateWarehouse(&warehouse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create warehouse",
			"details": err.Error(),
		})
		return
	}

	// Return the created warehouse as a response
	ctx.JSON(http.StatusCreated, gin.H{"data": warehouse})
}

func (c *WarehouseController) UpdateWarehouse(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	warehouse, err := c.warehouseService.GetWarehouseByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve warehouse", "details": err.Error()})
		return
	}
	// var warehouse models.Warehouse
	if err := ctx.ShouldBindJSON(&warehouse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	if err := c.warehouseService.UpdateWarehouse(warehouse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update warehouse", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": warehouse})
}

func (c *WarehouseController) DeleteWarehouse(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	wareHouse, err := c.warehouseService.GetWarehouseByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve warehouse", "details": err.Error()})
		return
	}

	if err := c.warehouseService.DeleteWarehouse(wareHouse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete warehouse", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Warehouse deleted successfully"})
}
