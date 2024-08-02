package mapping

import "time"

type User struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Email     string     `gorm:"column:email" json:"email"`
	Password  string     `gorm:"column:password" json:"password"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName
func (User) TableName() string {
	return "users"
}
