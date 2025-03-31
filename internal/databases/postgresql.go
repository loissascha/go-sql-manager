package databases

import "fmt"

type PostgreSQL struct {
}

func (p PostgreSQL) Connect() error {
	fmt.Println("Connected to PostgreSQL")
	return nil
}
