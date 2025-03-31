package databases

import (
	"database/sql"
	"fmt"
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
	fmt.Println("Connected to PostgreSQL")
	return nil, nil
}
