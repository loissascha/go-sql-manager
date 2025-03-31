package databases

import "fmt"

type MySQL struct {
}

func (m MySQL) Connect() error {
	fmt.Println("Connected to MySQL")
	return nil
}
