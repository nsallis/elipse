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
	Processor ProcessorPlugin
}

// ProcessorPlugin is the type we use to ensure the symbol returned from the plugin implements ProcessEntity
type ProcessorPlugin interface {
	Process(string) (string, error)
}

// GetNodeType get the type of node
func (n *GoProcessorNode) GetNodeType() string {
	return "GoProcessor"
}

// ToString get the string version of this node
func (n *GoProcessorNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutputs())
}

// Setup checks to make sure userCode exists. If it does, we run buildStringToSymbol, which
// writed the string to a temp file, compiles it as a plugin, and returns the symbol (with the Process function attached).
func (n *GoProcessorNode) Setup() { // TODO eventually just inject the code inside of the process string that is used as a template
	if userCode, exists := n.Config["userCode"]; exists {
		if !exists {
			log.Error("Required config `userCode` not present for node "+n.UUID, nil)
			panic("Required config `userCode` not present for node")
		}
		n.Processor = n.buildStringToSymbol(userCode)
	}
}

// Process run the worker. Uses the plugin we compiled in Setup, and runs it against incoming documents
func (n *GoProcessorNode) Process() {
	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				break
			}
		case document := <-n.InputChannel:
			output, err := n.Processor.Process(string(document.Value))
			if err != nil {
				log.Error("Could not process a document for node "+n.UUID, err)
				break
			}
			n.OutputChannels[0] <- n.buildDocumentFromValue(output, document)

		}
	}
}

func (n *GoProcessorNode) buildDocumentFromValue(value string, doc Document) Document {
	newDoc := doc
	newDoc.Value = []byte(value)
	return newDoc
}

// saves the passed string as a go file, then compiles it as a plugin. Finally, we check that the plugin defined a type with
// a `Process` function on it that taked and returns the correct types (see ProcessorPlugin interface for correct types)
func (n *GoProcessorNode) buildStringToSymbol(entity string) ProcessorPlugin {
	pluginObjectName := "processor_plugin_" + n.UUID
	pluginPath := "./.tmp/processorPlugins/" + pluginObjectName + ".go"
	sharedObjectPath := "./.tmp/processorPlugins/" + pluginObjectName + ".so"
	err := ioutil.WriteFile(pluginPath, []byte(entity), 0644)
	if err != nil {
		log.Error("There was an error writing to the plugin's temporary directory.", err)
	}
	out, err := exec.Command("go", "build", "-buildmode=plugin", "-o", sharedObjectPath, pluginPath).CombinedOutput()
	if err != nil {
		log.Error("Failed to compile user-provided processor code.\n"+string(out), nil)
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
