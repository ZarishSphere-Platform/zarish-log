package models

import "time"

// Shipment represents a logistics shipment
type Shipment struct {
	BaseModel

	TrackingNumber string `gorm:"size:100;uniqueIndex;not null" json:"tracking_number"`
	Status         string `gorm:"size:50;default:'pending'" json:"status"` // pending, in_transit, delivered, cancelled
	Type           string `gorm:"size:100" json:"type"`                    // medical_supplies, equipment, documents

	Origin      string `gorm:"size:500" json:"origin"`
	Destination string `gorm:"size:500" json:"destination"`

	ShippedAt   *time.Time `json:"shipped_at,omitempty"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`

	CarrierID *uint `gorm:"index" json:"carrier_id,omitempty"`
}

// Inventory represents inventory items
type Inventory struct {
	BaseModel

	ItemCode    string `gorm:"size:100;uniqueIndex;not null" json:"item_code"`
	Name        string `gorm:"size:255;not null;index" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Category    string `gorm:"size:100" json:"category"`

	Quantity    int    `gorm:"default:0" json:"quantity"`
	MinQuantity int    `gorm:"default:0" json:"min_quantity"`
	Unit        string `gorm:"size:50" json:"unit"`

	LocationID *uint `gorm:"index" json:"location_id,omitempty"`
}

// TableName overrides
func (Shipment) TableName() string {
	return "shipments"
}

func (Inventory) TableName() string {
	return "inventory_items"
}
