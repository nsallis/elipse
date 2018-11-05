package workers

import (
	"fmt"
	"github.com/nsallis/elipse/log"
	"io/ioutil"
	"os/exec"
	"plugin"
)

type GoProcessorNode struct {
	BaseNode
	UserCode ProcessorPlugin
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
	if userCode, exists := n.Config["userCode"]; exists {
		if !exists {
			log.Error("Required config `userCode` not present for node "+n.UUID, nil)
			panic("Required config `userCode` not present for node")
		}
		n.UserCode = n.buildStringToSymbol(userCode)
	}
}

func (n *GoProcessorNode) Process() {
	// TODO run n.UserCode.Process with the Document's value and pass on it's output
}

func (n *GoProcessorNode) buildStringToSymbol(entity string) ProcessorPlugin {
	pluginObjectName := "processor_plugin_" + n.UUID
	pluginPath := "./tmp/processorPlugins/" + pluginObjectName + ".go"
	sharedObjectPath := "./tmp/processorPlugins" + pluginObjectName + ".so"
	err := ioutil.WriteFile(pluginPath, []byte(entity), 0644)
	if err != nil {
		log.Error("There was an error writing to the plugin's temporary directory.", err)
	}
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", sharedObjectPath, pluginPath)
	err = cmd.Run()
	if err != nil {
		log.Error("Failed to compile user-provided processor code.", err)
		panic("Failed to compile user-provided processor code")
	}
	plug, err := plugin.Open(sharedObjectPath)
	if err != nil {
		log.Error("Cannot access shared object for user-provided processor code", err)
		panic("Cannot access shared object for user-provided processor code")
	}
	processFunc, err := plug.Lookup("Plugin")
	if err != nil {
		log.Error("Flailed to find `Plugin` export in user-provided code for GoProcessor", err)
		panic("Flailed to find `Plugin` export in user-provided code for GoProcessor")
	}
	var customProcessor ProcessorPlugin
	customProcessor, ok := processFunc.(ProcessorPlugin)
	if !ok {
		log.Error("User-provided processor code does not implement ProcessorPlugin interface.", err)
	}
	return customProcessor
}
