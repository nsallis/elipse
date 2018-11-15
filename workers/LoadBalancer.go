package workers

import (
	"fmt"
	"github.com/nsallis/elipse/log"
)

// LoadBalancerNode passes a given document to the output channel that is emptiest
// Requires `outputs` with at least one node
type LoadBalancerNode struct {
	BaseNode
}

// GetNodeType get the type of node
func (n *LoadBalancerNode) GetNodeType() string {
	return "LoadBalancer"
}

// ToString get the string version of this node
func (n *LoadBalancerNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutputs())
}

// Setup make any config updates before processing
func (n *LoadBalancerNode) Setup() {
	if len(n.OutputChannels) <= 0 {
		log.Error("Required config `outputs` not present for node "+n.UUID, nil)
		panic("Required config `outputs` not present for node")
	}
}

// Process run the worker. Passes the document to the emptiest channel
func (n *LoadBalancerNode) Process() {
	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				break
			}
		case document := <-n.InputChannel:
			var minChanLength int
			var emptiestChannel chan Document
			for i, outChan := range n.OutputChannels {
				if i == 0 {
					minChanLength = len(outChan)
					emptiestChannel = outChan
				} else {
					if len(outChan) < minChanLength {
						emptiestChannel = outChan
					} else {
						emptiestChannel = emptiestChannel // avoid no use error
					}

				}
			}
			emptiestChannel <- document
			break
		}
	}
}
