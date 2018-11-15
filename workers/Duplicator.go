package workers

import (
	"fmt"
	"github.com/nsallis/elipse/log"
)

// DuplicatorNode passes a given document to all output channels, duplicating the document
type DuplicatorNode struct {
	BaseNode
}

// GetNodeType get the type of node
func (n *DuplicatorNode) GetNodeType() string {
	return "Duplicator"
}

// ToString get the string version of this node
func (n *DuplicatorNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutputs())
}

// Setup make any config updates before processing
func (n *DuplicatorNode) Setup() {
	if len(n.OutputChannels) <= 0 {
		log.Error("Required config `outputs` not present for node "+n.UUID, nil)
		panic("Required config `outputs` not present for node")
	}
}

// Process run the worker. Passes the document to the emptiest channel
func (n *DuplicatorNode) Process() {
	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				break
			}
		case document := <-n.InputChannel:
			for _, outChan := range n.OutputChannels {
				docClone := document // clone document so we don't change it from another node
				outChan <- docClone
			}
			break
		}
	}
}
