package workers

import (
	"fmt"
	"log"
)

type GoProcessorNode struct {
	BaseNode
}

// ProcessorPlugin is the type we use to ensure the symbol returned from the plugin implements ProcessEntity
type ProcessorPlugin interface {
	Process(string) string
}

// GetNodeType get the type of node
func (n *GoProcessorNode) GetNodeType() string {
	return "GoProcessor"
}

// ToString get the string version of this node
func (n *GoProcessorNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}

func (n *GoProcessorNode) Setup() {

}

func (n *GoProcessorNode) Process() {

}

func (n *GoProcessorNode) buildStringToSymbol(entity string) {
	pluginObjectName := "processor_plugin_" + n.UUID + ".go"
	pluginPath := "./tmp/processorPlugins/" + pluginObjectName
	ioutil.WriteFile(pluginPath, []byte(processingInstructions), 0644)
}
