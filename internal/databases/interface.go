package databases

type Connector interface {
	Connect() error
}
