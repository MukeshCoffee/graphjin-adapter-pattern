package target

// SQLExecutor is the target interface for executing SQL commands
type SQLExecutor interface {
	ExecuteSQL() string
}
