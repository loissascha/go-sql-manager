package databases

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	connectionString string
}

func (p *PostgreSQL) SetConnectionString(con string) error {
	if con == "" {
		return fmt.Errorf("Can't set empty connection string!")
	}
	p.connectionString = con
	return nil
}

func (p *PostgreSQL) Connect() (*sql.DB, error) {
	if strings.TrimSpace(p.connectionString) == "" {
		return nil, fmt.Errorf("No connection string provided!")
	}
	db, err := sql.Open("postgres", p.connectionString)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")
	return db, nil
}

func (p *PostgreSQL) ListDatabases(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT datname FROM pg_database WHERE datistemplate = false")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var list []string
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return list, err
		}
		list = append(list, dbName)
	}
	return list, nil
}

func (p *PostgreSQL) ListTables(db *sql.DB, databaseName string) ([]string, error) {
	query := `SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname NOT IN ('pg_catalog', 'information_schema')`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	return tables, nil
}
