package adapter

import (
	"testing"
)

// TestAdapter_Request tests the Request method of the Adapter
func TestAdapter_Request(t *testing.T) {
	// Create an instance of Adaptee
	adaptee := &Adaptee{}

	// Pass the Adaptee instance to the Adapter
	adapter := &Adapter{Adaptee: adaptee}

	// Define the expected output
	expected := "Adaptee's specific request"

	// Get the actual output from the Adapter
	result := adapter.Request()

	// Compare the expected and actual outputs
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
