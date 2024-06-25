package port

// Closable this interface represent connection that can be closed like db connection
type Closable interface {
	Close() error
}
