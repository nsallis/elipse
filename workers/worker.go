package workers

type BaseNode struct {
  err error
}

type Configuration map[string]interface{} // should allow for any data scructure. Maybe this needs to be a map of string and interface?

type Node interface {
  Input(Document, Configuration) // will eventually take a document
  Process()
  Output() Document // should always return a document
  Error() (NodeError) // should always return an error or nil
}
