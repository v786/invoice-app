package routes

import (
	"jahangir/invoice/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)
	r.GET("/orders", controllers.FindOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/merchants", controllers.FindMerchants)
	r.POST("/merchants", controllers.CreateMerchant)
	return r
}
