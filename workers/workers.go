package workers

import "fmt"

// WorkerConfig configuration for all workers.
// Used in building nodes from json
// TODO should probably be moved to spawner
type WorkerConfig struct {
	UUID     string
	NodeType string
	Config   map[string]string
	Outputs  []string
	Errors   []string
}

// Node defines interface with methods all nodes should have
// use `var _ Node = DFINode{}` to check that your node implements
// Node
type Node interface {
	Setup() // will eventually take a document
	Process()
	SetUUID(string)              // setter
	SetConfig(map[string]string) // setter
	SetInput(chan Document)
	SetOutput(chan Document)
	SetError(chan error)
	SetControl(chan string)
	GetUUID() string
	GetConfig() map[string]string
	GetInput() chan Document
	GetOutput() chan Document
	GetError() chan error
	GetControl() chan string
	GetNodeType() string
	ToString() string
	// Output() Document // should always return a document
	// Error() (NodeError) // should always return an error or nil
}

// NodeStruct defined all of the values on a node. Every type of
// Node is of type NodeStruct
type BaseNode struct {
	UUID           string // make this a UUID eventually
	InputChannel   chan Document
	OutputChannel  chan Document
	ErrorChannel   chan error
	ControlChannel chan string // to "exit", "pause" etc.
	Config         map[string]string
}

func (n *BaseNode) Setup() {
	panic("not implemented")
}

func (n *BaseNode) Process() {
	panic("not implemented")
}

func (n *BaseNode) SetUUID(uuid string) {
	panic("not implemented")
}

func (n *BaseNode) SetConfig(config map[string]string) {
	panic("not implemented")
}

func (n *BaseNode) SetInput(input chan Document) {
	n.InputChannel = input
}

func (n *BaseNode) SetOutput(input chan Document) {
	n.OutputChannel = input
}

func (n *BaseNode) SetError(err chan error) {
	n.ErrorChannel = err
}

func (n *BaseNode) SetControl(control chan string) {
	n.ControlChannel = control
}

func (n *BaseNode) GetUUID() string {
	return n.UUID
}

func (n *BaseNode) GetConfig() map[string]string {
	return n.Config
}

func (n *BaseNode) GetInput() chan Document {
	return n.InputChannel
}

func (n *BaseNode) GetOutput() chan Document {
	return n.OutputChannel
}

func (n *BaseNode) GetError() chan error {
	return n.ErrorChannel
}

func (n *BaseNode) GetControl() chan string {
	return n.ControlChannel
}

func (n *BaseNode) GetNodeType() string {
	return "BaseNode"
}

func (n *BaseNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}
