package models
import "time"

type Assignment struct {
    ID          string     `gorm:"primaryKey;type:varchar(36)" json:"id"`
    AssetID     string     `gorm:"not null" json:"asset_id"`
    EmployeeID  string     `gorm:"not null" json:"employee_id"`
    AssignedAt  time.Time  `json:"assigned_at"`
    ReturnedAt  *time.Time `json:"returned_at"`
    AssignedBy  string     `json:"assigned_by"`
    Notes       string     `json:"notes"`
}
