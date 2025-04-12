package configs

import (
	"encoding/json"
	"go-sql-manager/internal/filesystem"
	"os"
)

type DatabaseType int

const (
	DATABASE_MYSQL DatabaseType = iota
	DATABASE_POSTGRES
)

type DatabaseConfig struct {
	User     string
	Host     string
	Port     string
	Password string
	Type     DatabaseType
}

type Database struct {
	initialized bool
	data        []DatabaseConfig
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) GetDatabaseConfigs() []DatabaseConfig {
	if d.initialized {
		return d.data
	}
	d.data = []DatabaseConfig{}
	filepath, err := filesystem.GetSavePath("databases.json")
	if err != nil {
		panic(err)
	}
	if !filesystem.FileExists(filepath) {
		return d.data
	}

	file, err := os.Open(filepath)
	if err != nil {
		panic("Can't open database.json file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&d.data); err != nil {
		panic("Can't decode database.json file")
	}
	return d.data
}
