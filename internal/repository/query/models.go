package query

import (
	"database/sql"
)

type Health struct {
	ID sql.NullInt32 `json:"id"`
}
