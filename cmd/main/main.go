package main

import (
	"fmt"
	"graphjin-adapter-pattern/internal/adaptee"
	"graphjin-adapter-pattern/internal/adapter"
	"graphjin-adapter-pattern/internal/target"
)

func main() {
	// First, demonstrate the old Adapter pattern
	fmt.Println("Demonstrating the old Adapter pattern:")

	// Create an instance of the previous Adaptee
	oldAdaptee := &adapter.Adaptee{}

	// Pass the old Adaptee instance to the old Adapter
	oldAdapter := &adapter.Adapter{Adaptee: oldAdaptee}

	// Use the old Adapter as the client would expect
	fmt.Println("Client:", oldAdapter.Request())

	fmt.Println() // Add a newline for better output separation

	// Now, demonstrate the new scenario with GraphQL to SQL conversion
	fmt.Println("Demonstrating the new GraphQL to SQL Adapter pattern:")

	// Create instances of the new adaptees
	getUserQuery := &adaptee.GetUserQuery{UserID: 302}
	createUserQuery := &adaptee.CreateUserQuery{UserName: "Mukesh Swain", UserEmail: "mukeshswain@coffeebeans.io"}

	// Create adapters for each new adaptee
	getUserAdapter := adapter.NewGetUserAdapter(getUserQuery)
	createUserAdapter := adapter.NewCreateUserAdapter(createUserQuery)

	// Create a slice of SQL executors
	executors := []target.SQLExecutor{getUserAdapter, createUserAdapter}

	// Execute SQL commands using both new adapters
	for _, executor := range executors {
		fmt.Println(executor.ExecuteSQL())
		fmt.Println() // Add a newline for better output separation
	}
}
