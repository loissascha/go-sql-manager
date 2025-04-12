package configs

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
	var result = []DatabaseConfig{}
	return result
}
