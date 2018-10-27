package workers

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
	// Output() Document // should always return a document
	// Error() (NodeError) // should always return an error or nil
}

// NodeStruct defined all of the values on a node. Every type of
// Node is of type NodeStruct
type NodeStruct struct {
	UUID           string // make this a UUID eventually
	InputChannel   chan Document
	OutputChannel  chan Document
	ErrorChannel   chan error
	ControlChannel chan string // to "exit", "pause" etc.
	Config         map[string]string
}
