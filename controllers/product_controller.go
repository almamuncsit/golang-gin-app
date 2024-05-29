package controllers

import (
    "github.com/gin-gonic/gin"
    "gin-app/database"
    "gin-app/models"
    "net/http"
)

// GetAllProducts retrieves all products from the database.
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
    var products []models.Product
    database.DB.Find(&products)
    c.JSON(http.StatusOK, products)
}

// GetProduct retrieves a product by its ID from the database.
// @Summary Get a product
// @Description Get a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} gin.H{"error": "Product not found"}
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")

    if err := database.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product in the database.
// @Summary Create a product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Product
// @Failure 400 {object} gin.H{"error": "Bad request"}
// @Router /products [post]
func CreateProduct(c *gin.Context) {
    var input models.Product

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&input)
    c.JSON(http.StatusOK, input)
}

// UpdateProduct updates an existing product in the database.
// @Summary Update a product
// @Description Update an existing product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Product
// @Failure 400 {object} gin.H{"error": "Bad request"}
// @Failure 404 {object} gin.H{"error": "Product not found"}
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")

    if err := database.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Model(&product).Updates(input)
    c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product from the database.
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} gin.H{"message": "Product deleted"}
// @Failure 404 {object} gin.H{"error": "Product not found"}
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")

    if err := database.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    database.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
