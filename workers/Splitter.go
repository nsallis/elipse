package workers

import (
	"bytes"
	"fmt"
	"github.com/nsallis/elipse/log"
)

// SplitterNode splits a document's value by a delimiter and outputs each split value
// as a separate document
// Takes config values:
//  - delimiter - the character(s) to split the value by
type SplitterNode struct {
	BaseNode
}

// GetNodeType get the type of node
func (n *SplitterNode) GetNodeType() string {
	return "SplitterNode"
}

// ToString get the string version of this node
func (n *SplitterNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}

// Setup make any config updates before processing
func (n *SplitterNode) Setup() {
	var exists bool
	if _, exists = n.Config["delimiter"]; !exists {
		log.Error("Required config `delimiter` not present in node"+n.UUID, nil)
		panic("Required config `delimiter` not present for node")
	}
}

// Process run the worker. Splits the value of a document based on a delimiter
// Writes these values into new documents, and outputs them
func (n *SplitterNode) Process() {
	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				break
			}
		case document := <-n.InputChannel:
			splitValues := bytes.Split(document.Value, []byte(n.Config["delimiter"]))
			for _, value := range splitValues {
				n.OutputChannel <- n.createDocFromString(value, document)
			}
		}
	}
}

func (n *SplitterNode) createDocFromString(value []byte, doc Document) Document {
	return Document{Value: value, Errors: doc.Errors, Source: doc.Source, SourceType: doc.SourceType, SourcePermissions: doc.SourcePermissions}
}
