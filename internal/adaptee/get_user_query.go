package adaptee

import "fmt"

// GetUserQuery represents a GraphQL query to get a user
type GetUserQuery struct {
	UserID int
}

// GenerateGraphQL returns the original GraphQL query
func (q *GetUserQuery) GenerateGraphQL() string {
	return fmt.Sprintf("{ user(id: %d) { id, name, email } }", q.UserID)
}

// GenerateSQL generates an SQL command from the GraphQL query
func (q *GetUserQuery) GenerateSQL() string {
	return fmt.Sprintf("SELECT * FROM users WHERE id = %d;", q.UserID)
}
