package query

import (
	"database/sql"
	"time"
)

type Stores struct {
	ID         sql.NullString `json:"id" gorm:"id"`
	StoreName  sql.NullString `json:"store_name" gorm:"store_name"`
	StorePhoto sql.NullString `json:"store_photo" gorm:"store_photo"`
	CreatedAt  time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"updated_at"`
}
