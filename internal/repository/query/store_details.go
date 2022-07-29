package query

import (
	"database/sql"
	"time"
)

type StoreDetails struct {
	ID           sql.NullString `json:"id" gorm:"id"`
	StoreID      sql.NullString `json:"store_id" gorm:"store_id"`
	StoreName    sql.NullString `json:"store_name" gorm:"store_name"`
	Owner        sql.NullString `json:"owner" gorm:"owner"`
	StorePhone   sql.NullString `json:"store_phone" gorm:"store_phone"`
	StorePhoto   sql.NullString `json:"store_photo" gorm:"store_photo"`
	StoreAddress sql.NullString `json:"store_address" gorm:"store_address"`
	CreatedAt    time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"updated_at"`
}
