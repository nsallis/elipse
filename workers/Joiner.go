package workers

import (
	"fmt"
	"github.com/nsallis/elipse/log"
	"strings"
)

type JoinerNode struct {
	BaseNode
	InProgresDocuments map[string][]string // holds document values that are in progress
}

// GetNodeType get the type of node
func (n *JoinerNode) GetNodeType() string {
	return "Joiner"
}

// ToString get the string version of this node
func (n *JoinerNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}

// Setup makes updates to config before processing
// config options:
//  - delimiter -> how to join the lines
func (n *JoinerNode) Setup() {
	_, exists := n.Config["delimiter"]
	if !exists {
		log.Error("Required config `delimiter` not present for node "+n.UUID, nil)
		panic("Required config `delimiter` not present for node")
	}
	n.InProgresDocuments = make(map[string][]string)
}

func (n *JoinerNode) Process() {
	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				break
			}
		case document := <-n.InputChannel:
			if _, exists := n.InProgresDocuments[document.Source]; !exists {
				n.InProgresDocuments[document.Source] = []string{string(document.Value)}
			} else {
				n.InProgresDocuments[document.Source] = append(n.InProgresDocuments[document.Source], string(document.Value))
			}
			// log.Debug("total fragments: " + string(document.TotalFragments))
			var currentLength int
			currentLength = len(n.InProgresDocuments[document.Source])
			if currentLength%1000 == 0 {
				log.Debug(fmt.Sprintf("At fragment: %d", currentLength))
			}
			if currentLength >= document.TotalFragments {
				log.Debug("Got all fragments. Joining them now")
				newVal := strings.Join(n.InProgresDocuments[document.Source][:], n.Config["delimiter"])
				log.Debug("writing to output channel")
				n.OutputChannel <- n.createDocFromValue(newVal, document)
				delete(n.InProgresDocuments, document.Source)
			}
		}
	}
}

func (n *JoinerNode) createDocFromValue(value string, doc Document) Document {
	newDoc := doc
	newDoc.Value = []byte(value)
	return newDoc
}
