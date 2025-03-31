package databases

type Database interface {
	SetConnectionString(string) error
	Connect() error
}
