package databases

import "database/sql"

type Database interface {
	SetConnectionString(string) error
	Connect() (*sql.DB, error)
	ListDatabases(*sql.DB) ([]string, error)
	ListTables(db *sql.DB, databaseName string) ([]string, error)
	ListTable(db *sql.DB, databaseName string, tableName string) error
}
