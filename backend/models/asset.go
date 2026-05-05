package models
import ("time"; "github.com/google/uuid"; "gorm.io/gorm")

type Asset struct {
    ID              string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
    AssetCode       string    `gorm:"uniqueIndex;not null" json:"asset_code"`
    Name            string    `gorm:"not null" json:"name"`
    Category        string    `json:"category"`
    SerialNumber    string    `json:"serial_number"`
    PurchaseDate    time.Time `json:"purchase_date"`
    PurchaseValue   float64   `json:"purchase_value"`
    CurrentValue    float64   `json:"current_value"`
    Location        string    `json:"location"`
    Status          string    `gorm:"default:ACTIVE" json:"status"`
    PurchaseOrderID string    `json:"purchase_order_id"`
    Notes           string    `json:"notes"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
func (a *Asset) BeforeCreate(tx *gorm.DB) error {
    a.ID = uuid.NewString(); return nil
}