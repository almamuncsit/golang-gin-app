package routes

import (
    "github.com/gin-gonic/gin"
    "gin-app/controllers"
)

func UserRoutes(router *gin.RouterGroup) {
    userGroup := router.Group("/users")
    {
        userGroup.GET("/", controllers.GetAllUsers)
        userGroup.GET("/:id", controllers.GetUser)
        userGroup.POST("/", controllers.CreateUser)
        userGroup.PUT("/:id", controllers.UpdateUser)
        userGroup.DELETE("/:id", controllers.DeleteUser)
    }
}
