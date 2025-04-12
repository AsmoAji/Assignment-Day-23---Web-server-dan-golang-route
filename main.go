
package main

import (
    "inventory-app/config"
    "inventory-app/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    config.ConnectDatabase()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
