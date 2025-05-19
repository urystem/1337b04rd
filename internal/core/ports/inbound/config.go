package inbound

type Configs interface {
	GetHostPort() (string, error)
}
