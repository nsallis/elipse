package workers


type BaseNode struct {
  err error
}

type Configuration map[string]interface{} // should allow for any data scructure. Maybe this needs to be a map of string and interface?

type Node interface {
  Setup() // will eventually take a document
  Process()
  // Output() Document // should always return a document
  // Error() (NodeError) // should always return an error or nil
}

type NodeStruct struct {
  UUID string // make this a UUID eventually
  InputChannel chan Document
  OutputChannel chan Document
  ErrorChannel chan error
  ControlChannel chan string // to "exit", "pause" etc.
  Config Configuration
}
