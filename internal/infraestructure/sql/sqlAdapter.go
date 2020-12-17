package sql

import "database/sql"

type Connection interface {
	Get() *sql.DB
}
