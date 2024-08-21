package adapter

// Adaptee is the existing class with an incompatible interface
type Adaptee struct{}

// SpecificRequest is a method that the Adaptee class implements
func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's specific request"
}
