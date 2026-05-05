package models
import "time"

type Maintenance struct {
    ID            string     `gorm:"primaryKey;type:varchar(36)" json:"id"`
    AssetID       string     `gorm:"not null" json:"asset_id"`
    ScheduledDate time.Time  `json:"scheduled_date"`
    CompletedDate *time.Time `json:"completed_date"`
    Cost          float64    `json:"cost"`
    Description   string     `json:"description"`
    Status        string     `gorm:"default:SCHEDULED" json:"status"`
}