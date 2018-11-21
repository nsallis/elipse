package workers

import "fmt"

// Node defines interface with methods all nodes should have
// use `var _ Node = DFINode{}` to check that your node implements
// Node
type Node interface {
	Setup() // will eventually take a document
	Process()
	SetUUID(string)              // setter
	SetConfig(map[string]string) // setter
	SetInput(chan Document)
	AddOutputNode(string)
	DelOutputNode(string)
	AddOutput(chan Document)
	DelOutputs()
	SetError(chan error)
	SetControl(chan string)
	GetUUID() string
	GetConfig() map[string]string
	GetInput() chan Document
	GetOutputs() []chan Document
	GetError() chan error
	GetControl() chan string
	GetNodeType() string
	GetOutputNodes() []string
	ToString() string
	// Output() Document // should always return a document
	// Error() (NodeError) // should always return an error or nil
}

// NodeStruct defined all of the values on a node. Every type of
// Node is of type NodeStruct
type BaseNode struct {
	UUID           string          // make this a UUID eventually
	InputChannel   chan Document   `json:"-"`
	OutputChannels []chan Document `json:"-"`
	ErrorChannel   chan error      `json:"-"`
	ControlChannel chan string     `json:"-"` // to "exit", "pause" etc.
	Config         map[string]string
	OutputNodes    []string `json:"Outputs"` // used for serializing connections between nodes
}

func (n *BaseNode) Setup() {
	panic("not implemented")
}

func (n *BaseNode) Process() {
	panic("not implemented")
}

func (n *BaseNode) SetUUID(uuid string) {
	n.UUID = uuid
}

func (n *BaseNode) SetConfig(config map[string]string) {
	n.Config = config
}

func (n *BaseNode) SetInput(input chan Document) {
	n.InputChannel = input
}

func (n *BaseNode) AddOutputNode(uuid string) {
	n.OutputNodes = append(n.OutputNodes, uuid)
}

func (n *BaseNode) DelOutputNode(uuid string) {
	var newOutputs []string
	for _, item := range n.OutputNodes {
		if item != uuid {
			newOutputs = append(newOutputs, uuid)
		}
	}
	n.OutputNodes = newOutputs
}

func (n *BaseNode) AddOutput(input chan Document) {
	n.OutputChannels = append(n.OutputChannels, input)
}

func (n *BaseNode) DelOutputs() {
	n.OutputChannels = []chan Document{}
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

func (n *BaseNode) GetOutputs() []chan Document {
	return n.OutputChannels
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

func (n *BaseNode) GetOutputNodes() []string {
	return n.OutputNodes
}

func (n *BaseNode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutputs())
}
