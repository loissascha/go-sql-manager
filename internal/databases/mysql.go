package databases

import "fmt"

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
	fmt.Println("Connected to MySQL")
	return nil
}
