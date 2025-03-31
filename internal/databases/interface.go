package databases

type Database interface {
	Connect() error
}
