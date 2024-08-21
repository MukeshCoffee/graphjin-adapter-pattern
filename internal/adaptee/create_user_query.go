package adaptee

import "fmt"

// CreateUserQuery represents a GraphQL query to create a user
type CreateUserQuery struct {
	UserName  string
	UserEmail string
}

// GenerateGraphQL returns the original GraphQL query
func (q *CreateUserQuery) GenerateGraphQL() string {
	return fmt.Sprintf("mutation { createUser(name: \"%s\", email: \"%s\") { id, name, email } }", q.UserName, q.UserEmail)
}

// GenerateSQL generates an SQL command from the GraphQL query
func (q *CreateUserQuery) GenerateSQL() string {
	return fmt.Sprintf("INSERT INTO users (name, email) VALUES ('%s', '%s');", q.UserName, q.UserEmail)
}
