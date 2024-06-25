package port

// Server is an interface that represents a server.
type Server interface {
	Start() error
	Stop() error
}
