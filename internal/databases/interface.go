package databases

import "database/sql"

type Database interface {
	SetConnectionString(string) error
	Connect() (*sql.DB, error)
}
