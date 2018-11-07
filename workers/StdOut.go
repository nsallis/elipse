// STD out logger. Logs input documents to stdout
package workers

import (
	"fmt"
	"github.com/nsallis/elipse/log"
	"time"
)

// StdOutNode prints documents to standard out (mostly for debugging)
type StdOutNode struct {
	BaseNode
}

// GetNodeType gets the type of node
func (n *StdOutNode) GetNodeType() string {
	return "StdOut"
}

// ToString returns the string version of this node
func (n *StdOutNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}

// Setup runs any config updates needed before processing (non in this case)
func (n *StdOutNode) Setup() {

}

// Process runs the worker. Outputs any incoming documents to stdout
func (n *StdOutNode) Process() {
	defer close(n.OutputChannel)
	for {
		select {
		case inputDoc := <-n.InputChannel:
			docString, _ := inputDoc.ToString()
			fmt.Println("Got a document: " + docString)
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				close(n.OutputChannel)
				break
			}

		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}
