package main

import (
    "github.com/gin-gonic/gin"
    "gin-app/database"
    "gin-app/middleware"
    "gin-app/routes"
)

func main() {
    // Connect to the database
    database.ConnectDatabase()

    router := gin.Default()

    // Apply global middleware
    router.Use(middleware.LoggingMiddleware())

    // Apply authentication middleware to routes that need it
    authGroup := router.Group("/")
    authGroup.Use(middleware.AuthMiddleware())
    {
        routes.UserRoutes(authGroup)
        routes.ProductRoutes(authGroup)
    }

    // Start the server on port 8080
    router.Run(":8080")
}
