package server

import (
	"github.com/nsallis/elipse/log"
)

// Endpoints struct that all rpc handler endpoints are hung on
type Endpoints struct{}

// Handshake test endpoint to make sure our rpc server is running
func (e *Endpoints) Handshake(req Request, res *Response) error {
	log.Debug("Got handshake from rpc client")
	res.Payload = "Handshake successfull"
	return nil
}

func (e *Endpoints) GetWorkers(req Request, res *Response) error {
	log.Debug("Requested workers with payload: %v", req.Payload)
	return nil
}
