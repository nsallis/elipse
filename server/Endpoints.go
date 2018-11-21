package server

import (
	"encoding/json"
	// "github.com/Jeffail/gabs"
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
	// must pass Payload as \ escaped string!
	// jsonParsed, err := gabs.ParseJSON([]byte(req.Payload))
	// if err != nil {
	// 	log.Error("Unable to parse payload for rpc request", err)
	// }
	// log.Debug("Requested workers with payload: %v", jsonParsed)
	// foo := jsonParsed.Search("foo")
	// log.Debug("foo: %s", foo)
	json, err := json.Marshal(WorkersMap)
	if err != nil {
		log.Error("Could not convert workersMap to json", err)
	}
	res.Payload = string(json)
	return nil
}
