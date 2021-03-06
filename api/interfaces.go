package api

// InternalAPI is the interface our internal API uses
type InternalAPI interface {
	// Serialize transforms the data structure into a byte slice to be sent
	// via the websocket - this should be done as a serialized JSON string
	Serialize() []byte
}

// ExternalAPI is the interface for all our APIs
type ExternalAPI interface {
	// Get takes in an arbitrary argument, and builds the struct
	Get(interface{})

	// Transform takes the API data and turns it the data we are going to
	// send to the frontend
	Transform(interface{}) InternalAPI
}
