package adapter

// Adapter is the struct that adapts Adaptee to the Target interface
type Adapter struct {
	Adaptee *Adaptee
}

// Request is the method that makes Adaptee's interface compatible with Target's
func (a *Adapter) Request() string {
	// Here we are adapting the SpecificRequest method of Adaptee to be compatible with the Request method of Target
	return a.Adaptee.SpecificRequest()
}
