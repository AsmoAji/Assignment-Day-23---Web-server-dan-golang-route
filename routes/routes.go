
package routes

import (
    "github.com/gin-gonic/gin"
    "inventory-app/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/products", controllers.GetAllProducts)
    r.GET("/products/:id", controllers.GetProductByID)
    r.GET("/products/category/:category", controllers.GetProductByCategory)
    r.POST("/products", controllers.CreateProduct)
    r.PUT("/products/:id", controllers.UpdateProduct)
    r.DELETE("/products/:id", controllers.DeleteProduct)
    r.POST("/products/:id/upload", controllers.UploadImage)
    r.Static("/uploads", "./uploads")

    r.GET("/inventory/:product_id", controllers.GetInventory)
    r.PUT("/inventory/:product_id", controllers.UpdateInventory)

    r.POST("/orders", controllers.CreateOrder)
    r.GET("/orders/:id", controllers.GetOrderByID)

    r.GET("/report/stock", controllers.GetStockReport)
    r.GET("/report/orders", controllers.GetOrderReport)
}
