// STD out logger. Logs input documents to stdout
package workers

import (
	"fmt"
	"time"
)

type StdOutNode struct {
	BaseNode
}

func (n *StdOutNode) GetNodeType() string {
	return "StdOut"
}

func (n *StdOutNode) SetUUID(uuid string) {
	n.UUID = uuid
}

func (n *StdOutNode) SetConfig(config map[string]string) {
	n.Config = config
}

func (n *StdOutNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}

func (n *StdOutNode) Setup() {

}

func (n *StdOutNode) Process() {
	for {
		select {
		case inputDoc := <-n.InputChannel:
			docString, _ := inputDoc.ToString()
			fmt.Println("Got a document: " + docString)
		case command := <-n.ControlChannel:
			if command == "exit" {
				fmt.Println("exiting stdout...")
				close(n.OutputChannel)
				break
			}

		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}
