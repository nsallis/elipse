package server

import (
	"fmt"
	"github.com/nsallis/elipse/log"
	"github.com/nsallis/elipse/workers"
	"net"
	"net/rpc"
)

var WorkersMap map[string]workers.Node

type Server struct {
	WorkersMap map[string]workers.Node
}

func (s *Server) Initialize(workersMap map[string]workers.Node) {
	WorkersMap = workersMap
}

func (s *Server) Start(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error("There was an error starting the rpc server", err)
	}
	defer listener.Close()

	rpc.Register(&Endpoints{})
	rpc.Accept(listener)
}
