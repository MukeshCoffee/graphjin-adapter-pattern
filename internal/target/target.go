package target

// Target interface that the client expects
type Target interface {
	Request() string
}
