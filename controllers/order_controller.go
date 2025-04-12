package controllers

import (
	"inventory-app/config"
	"inventory-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inventory models.Inventory
	if err := config.DB.Where("product_id = ?", input.ProductID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventaris tidak ditemukan"})
		return
	}

	if inventory.Quantity < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stok tidak cukup"})
		return
	}

	order := models.Order{
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		OrderDate: time.Now(),
	}
	config.DB.Create(&order)

	inventory.Quantity -= input.Quantity
	config.DB.Save(&inventory)

	c.JSON(http.StatusCreated, order)
}

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, order)
}

type StockReport struct {
	Location   string
	TotalStock int
}

func GetStockReport(c *gin.Context) {
	var result []StockReport
	config.DB.Raw("SELECT location, SUM(quantity) as total_stock FROM inventories GROUP BY location").Scan(&result)
	c.JSON(http.StatusOK, result)
}

type OrderReport struct {
	ProductID    uint
	TotalOrdered int
}

func GetOrderReport(c *gin.Context) {
	var result []OrderReport
	config.DB.Raw("SELECT product_id, SUM(quantity) as total_ordered FROM orders GROUP BY product_id").Scan(&result)
	c.JSON(http.StatusOK, result)
}
