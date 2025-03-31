package databases

import (
	"database/sql"
	"fmt"
	"log"
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

func (m *MySQL) ListDatabases(db *sql.DB) {
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

	fmt.Println("Databases:")
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			log.Fatal(err)
		}
		fmt.Println("-", dbName)
	}
}
