// controllers/books.go

package controllers

import (
	// "bookc/models"
	"jahangir/invoice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateOrderInput struct {
	CreatedBy    string `json:"createdBy"`
	Amount       string `json:"amount"`
	OrderDetails string `json:"orderDetails`
	Mid          uint   `json:"mid`
}

type UpdateOrderInput struct {
	CreatedBy    string `json:"createdBy"`
	Amount       string `json:"amount"`
	OrderDetails string `json:"orderDetails`
}

// GET /orders
// Get all tasks
func FindOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var orders []models.Order
	db.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// POST /tasks
// Create new task
func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	amount, err := strconv.Atoi(input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var merchant models.Merchant
	db.First(&merchant, input.Mid)
	order := models.Order{
		CreatedBy:    input.CreatedBy,
		Amount:       amount,
		OrderDetails: input.OrderDetails,
		Mid:          merchant.ID,
	}
	db.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// // GET /tasks/:id
// // Find a task
// func FindTask(c *gin.Context) { // Get model if exist
// 	var task models.Task

// 	db := c.MustGet("db").(*gorm.DB)
// 	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": task})
// }

// // PATCH /tasks/:id
// // Update a task
// func UpdateTask(c *gin.Context) {

// 	db := c.MustGet("db").(*gorm.DB)
// 	// Get model if exist
// 	var task models.Task
// 	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input UpdateTaskInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	date := "2006-01-02"
// 	deadline, _ := time.Parse(date, input.Deadline)

// 	var updatedInput models.Task
// 	updatedInput.Deadline = deadline
// 	updatedInput.AssingedTo = input.AssingedTo
// 	updatedInput.Task = input.Task

// 	db.Model(&task).Updates(updatedInput)

// 	c.JSON(http.StatusOK, gin.H{"data": task})
// }

// // DELETE /tasks/:id
// // Delete a task
// func DeleteTask(c *gin.Context) {
// 	// Get model if exist
// 	db := c.MustGet("db").(*gorm.DB)
// 	var book models.Task
// 	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	db.Delete(&book)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
