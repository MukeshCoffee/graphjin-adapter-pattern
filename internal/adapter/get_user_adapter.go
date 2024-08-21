package adapter

import (
	"graphjin-adapter-pattern/internal/adaptee"
)

// GetUserAdapter adapts GetUserQuery to the SQLExecutor interface
type GetUserAdapter struct {
	query *adaptee.GetUserQuery
}

// NewGetUserAdapter creates a new GetUserAdapter
func NewGetUserAdapter(query *adaptee.GetUserQuery) *GetUserAdapter {
	return &GetUserAdapter{query: query}
}

// ExecuteSQL adapts the GenerateSQL method to the ExecuteSQL method
func (a *GetUserAdapter) ExecuteSQL() string {
	graphQL := a.query.GenerateGraphQL()
	sql := a.query.GenerateSQL()
	return "GraphQL Query: " + graphQL + "\nSQL Command: " + sql
}
