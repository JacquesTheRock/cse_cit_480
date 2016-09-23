package util

import (
	"database/sql"
	_ "github.com/lib/pq" //Required for Postgres
)

var Database *sql.DB
