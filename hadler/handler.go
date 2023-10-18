package hadler

import (
	"github.com/Asad2730/RDS-DB-Auth/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/", service.GetUsers)
	r.POST("/", service.CreateUser)
	r.PUT("/:email", service.UpdateUser)
}
