package databases

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"github.com/loissascha/go-logger/logger"
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
	db, err := sql.Open("postgres", p.connectionString+"postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	logger.Info(nil, "Connected to PostgreSQL")
	return db, nil
}

func (p *PostgreSQL) ListDatabases(db *sql.DB) ([]string, error) {
	logger.Info(nil, "SHOW DATABASES query")
	rows, err := db.Query("SELECT datname FROM pg_database WHERE datistemplate = false")
	if err != nil {
		logger.Error(err, "SHOW DATABASES error!")
		fmt.Println(err)
		return []string{}, err
	}
	defer rows.Close()

	var list []string
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			logger.Error(err, "SHOW DATABASES row scan error!")
			return list, err
		}
		list = append(list, dbName)
	}
	return list, nil
}

func (p *PostgreSQL) ListTables(dba *sql.DB, databaseName string) ([]string, error) {
	db, err := sql.Open("postgres", p.connectionString+databaseName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()
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
	fmt.Println(tables)
	return tables, nil
}

func (m *PostgreSQL) ListTable(dba *sql.DB, databaseName string, tableName string) error {
	db, err := sql.Open("postgres", m.connectionString+databaseName+"?sslmode=disable")
	if err != nil {
		logger.Error(err, "Error creating db connection for list table!")
		return err
	}
	defer db.Close()
	return nil
}
