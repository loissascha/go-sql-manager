package databases

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	connectionString string
}

func (m *MySQL) SetConnectionString(con string) error {
	if con == "" {
		return fmt.Errorf("Can't set empty connection string!")
	}
	m.connectionString = con
	return nil
}

func (m *MySQL) Connect() (*sql.DB, error) {
	if strings.TrimSpace(m.connectionString) == "" {
		return nil, fmt.Errorf("No connection string provided!")
	}
	db, err := sql.Open("mysql", m.connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MySQL")
	return db, nil
}

func (m *MySQL) ListDatabases(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	list := []string{}
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return list, err
		}
		list = append(list, dbName)
	}
	return list, nil
}

func (m *MySQL) ListTables(db *sql.DB, databaseName string) ([]string, error) {
	query := fmt.Sprintf("SHOW TABLES FROM `%s`", databaseName)
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

func (m *MySQL) ListTable(db *sql.DB, databaseName string, tableName string) error {
	return nil
}
