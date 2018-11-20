package server

// Request definition of rpc request
type Request struct {
	Payload string "json:Payload"
}

// Response definition of rpc response
type Response struct {
	Payload string
}
