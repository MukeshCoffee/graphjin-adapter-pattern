package adapter

import (
	"graphjin-adapter-pattern/internal/adaptee"
	"testing"
)

func TestCreateUserAdapter_ExecuteSQL(t *testing.T) {
	// Arrange
	query := &adaptee.CreateUserQuery{UserName: "Mukesh Swain", UserEmail: "mukeshswain@coffeebeans.io"}
	adapter := NewCreateUserAdapter(query)

	// Act
	result := adapter.ExecuteSQL()

	// Assert
	expected := "GraphQL Query: mutation { createUser(name: \"Mukesh Swain\", email: \"mukeshswain@coffeebeans.io\") { id, name, email } }\nSQL Command: INSERT INTO users (name, email) VALUES ('Mukesh Swain', 'mukeshswain@coffeebeans.io');"
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
