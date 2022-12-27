package src

import (
	"golangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group(("/users"))

	user.GET("/", service.GetUser)
	user.POST("/", service.PostUser)
	user.GET("/:id", service.FindByUserId)
	user.POST("/delete/:id", service.DeleteUser)
	user.POST("/modify/:id", service.PutUser)
}
