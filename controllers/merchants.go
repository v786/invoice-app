// controllers/books.go

package controllers

import (
	// "bookc/models"
	"jahangir/invoice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateMerchantInput struct {
	UserName    string `json:"userName"`
	TinNumber   string `json:"tinNumber"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type UpdateMerchantInput struct {
	UserName    string `json:"userName"`
	TinNumber   string `json:"tinNumber"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

// GET /orders
// Get all merchant
func FindMerchants(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var orders []models.Merchant
	db.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// POST /merchant
// Create new merchant
func CreateMerchant(c *gin.Context) {
	// Validate input
	var input CreateMerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create merchant
	merchant := models.Merchant{
		UserName:    input.UserName,
		TinNumber:   input.TinNumber,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&merchant)

	c.JSON(http.StatusOK, gin.H{"data": merchant})
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
