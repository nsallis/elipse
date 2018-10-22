package workers

// disk file input

import (
  "io/ioutil"
  // "path/filepath"
  "fmt"
  "github.com/fsnotify/fsnotify"
  "time"
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

func (n *DFINode) Process() {
  // TODO create document from file
  // TODO output via outputChannel

  var filepath string
  if str, ok := n.Config["filename"].(string); ok {
    filepath = str
  }

  if str, ok := n.Config["filename"].(string); ok {
    fileContents, err := ioutil.ReadFile(str)
    fmt.Println(string(fileContents), err)
  }
  // var watcher *fsnotify.Watcher
  watcher, _ := fsnotify.NewWatcher()
  // defer watcher.Close()
  watcher.Add(filepath)
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
    case event := <- watcher.Events:
      switch {
      case event.Op == fsnotify.Write:
        fmt.Println("wrote to the file")
        fileContents, _ := ioutil.ReadFile(event.Name)
        createDocFromNode(n, fileContents)
      case event.Op == fsnotify.Create:
        fmt.Println("created a file in a watched directory")
      default:

      }
      fmt.Println(event)
    default:
      time.Sleep(time.Second * 1)
    }
  }
}

func createDocFromNode(node *DFINode, fileContents []byte) Document {

  fmt.Println("file contents to create doc from: " + string(fileContents))
  var filepath string
  if str, ok := node.Config["filename"].(string); ok {
    filepath = str
  }
  return Document{Value: fileContents, Source: filepath, SourceType: "disk"}
}

