package mapping

import "time"

type Items struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Price     int        `gorm:"column:price" json:"price"`
	Quantity  int        `gorm:"column:quantity" json:"quantity"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName
func (Items) TableName() string {
	return "items"
}
