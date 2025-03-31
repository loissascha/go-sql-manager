package databases

import (
	"fmt"
	"strings"
)

type MySQL struct {
	connectionString string
}

func (m MySQL) SetConnectionString(con string) error {
	if con == "" {
		return fmt.Errorf("Can't set empty connection string!")
	}
	m.connectionString = con
	return nil
}

func (m MySQL) Connect() error {
	if strings.TrimSpace(m.connectionString) == "" {
		return fmt.Errorf("No connection string provided!")
	}
	fmt.Println("Connected to MySQL")
	return nil
}
