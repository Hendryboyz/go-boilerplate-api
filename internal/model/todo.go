package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID  `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	UserId      string     `gorm:"column:user_id;not null"`
	Description string     `gorm:"column:description;not null"`
	StartDate   *time.Time `gorm:"column:start_date;type:timestamp;not null"`
	EndDate     *time.Time `gorm:"column:end_date;type:timestamp;not null"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:timestamp"`
}
