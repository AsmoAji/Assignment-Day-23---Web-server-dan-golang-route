
package models

import "time"

type Order struct {
    ID        uint `gorm:"primaryKey"`
    ProductID uint
    Quantity  int
    OrderDate time.Time
}
