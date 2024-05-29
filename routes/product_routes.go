package routes

import (
    "github.com/gin-gonic/gin"
    "gin-app/controllers"
)

func ProductRoutes(router *gin.RouterGroup) {
    productGroup := router.Group("/products")
    {
        productGroup.GET("/", controllers.GetAllProducts)
        productGroup.GET("/:id", controllers.GetProduct)
        productGroup.POST("/", controllers.CreateProduct)
        productGroup.PUT("/:id", controllers.UpdateProduct)
        productGroup.DELETE("/:id", controllers.DeleteProduct)
    }
}
