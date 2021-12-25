package main

import (
	"jahangir/invoice/models"
	"jahangir/invoice/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Merchant{})

	r := routes.SetupRoutes(db)
	r.Run()
}

func migrate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Order{})
}
