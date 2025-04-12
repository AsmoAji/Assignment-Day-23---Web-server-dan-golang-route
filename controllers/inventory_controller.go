
package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "inventory-app/config"
    "inventory-app/models"
)

func GetInventory(c *gin.Context) {
    productID, err := strconv.Atoi(c.Param("product_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
        return
    }

    var inventory models.Inventory
    if err := config.DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Inventaris tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, inventory)
}

func UpdateInventory(c *gin.Context) {
    productID, err := strconv.Atoi(c.Param("product_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
        return
    }

    var inventory models.Inventory
    if err := config.DB.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Inventaris tidak ditemukan"})
        return
    }

    var updateData struct {
        Quantity int `json:"quantity"`
    }
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    inventory.Quantity += updateData.Quantity
    config.DB.Save(&inventory)
    c.JSON(http.StatusOK, inventory)
}
