package adapter

import (
	"graphjin-adapter-pattern/internal/adaptee"
	"testing"
)

func TestGetUserAdapter_ExecuteSQL(t *testing.T) {
	// Arrange
	query := &adaptee.GetUserQuery{UserID: 302}
	adapter := NewGetUserAdapter(query)

	// Act
	result := adapter.ExecuteSQL()

	// Assert
	expected := "GraphQL Query: { user(id: 302) { id, name, email } }\nSQL Command: SELECT * FROM users WHERE id = 302;"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
