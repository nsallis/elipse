package workers

// disk file input

import (
  "io/ioutil"
  // "path/filepath"
  "fmt"
)

type DFINode NodeStruct

func (n *DFINode) Setup() { // will eventually take a document
  // var _ Node = DFINode{}
  // TODO check if file path is valid
  // TODO get the absolute path of the file
  if str, ok := n.Config["filename"].(string); ok {
    n.Config["filename"] = str
  }
}

func (n *DFINode) Process() { // Just an empty document since this is an input node
  // inputChannel := n.InputChannel
  // TODO ignore input channel
  // TODO listen for file changes or run on current file if this is first time running
  // TODO create document from file
  // TODO output via outputChannel

  if str, ok := n.Config["filename"].(string); ok {
    fileContents, err := ioutil.ReadFile(str)
    fmt.Println(string(fileContents), err)
  }
  for{
    select {
    case command := <- n.ControlChannel:
      if(command == "exit"){
        close(n.InputChannel)
        close(n.OutputChannel)
        close(n.ControlChannel)
        close(n.ErrorChannel)
        break
      }
    default:
      // fmt.Println("Processing...")
      // TODO watch for file changes and process them
    }
  }

}
