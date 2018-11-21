package server

// Request definition of rpc request
type Request struct {
	Payload string
}

// Response definition of rpc response
type Response struct {
	Payload string
}

// Payload gets populated when rpc requests get parsed from string
type Payload map[string]interface{}
