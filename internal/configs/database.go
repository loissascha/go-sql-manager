package configs

import (
	"encoding/json"
	"fmt"
	"go-sql-manager/internal/filesystem"
	"os"
)

type DatabaseType int

const (
	DATABASE_MYSQL DatabaseType = iota
	DATABASE_POSTGRES
)

type DatabaseConfig struct {
	Id       string
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

func (d *Database) AddDatabaseConfig(config DatabaseConfig) error {
	d.data = append(d.data, config)
	err := d.SaveDatabaseConfigs()
	return err
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
		d.initialized = true
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
	d.initialized = true
	return d.data
}

func (d *Database) SaveDatabaseConfigs() error {
	if !d.initialized {
		return fmt.Errorf("Database not yet initialized. Please load data before trying to save.")
	}
	filepath, err := filesystem.GetSavePath("databases.json")
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(d.data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
