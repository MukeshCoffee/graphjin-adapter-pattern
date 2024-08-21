package adapter

import (
	"graphjin-adapter-pattern/internal/adaptee"
)

// CreateUserAdapter adapts CreateUserQuery to the SQLExecutor interface
type CreateUserAdapter struct {
	query *adaptee.CreateUserQuery
}

// NewCreateUserAdapter creates a new CreateUserAdapter
func NewCreateUserAdapter(query *adaptee.CreateUserQuery) *CreateUserAdapter {
	return &CreateUserAdapter{query: query}
}

// ExecuteSQL adapts the GenerateSQL method to the ExecuteSQL method
func (a *CreateUserAdapter) ExecuteSQL() string {
	graphQL := a.query.GenerateGraphQL()
	sql := a.query.GenerateSQL()
	return "GraphQL Query: " + graphQL + "\nSQL Command: " + sql
}
