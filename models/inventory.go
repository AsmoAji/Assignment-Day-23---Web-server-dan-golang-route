
package models

type Inventory struct {
    ID        uint `gorm:"primaryKey"`
    ProductID uint
    Quantity  int
    Location  string
}
