package main

import (
  "fmt"
  "github.com/nsallis/elipse/workers"
  "time"
  // "os"
)

func main() {
  inputChannel := make(chan workers.Document)
  outputChannel := make(chan workers.Document)
  controlChannel := make(chan string)
  errorChannel := make(chan error)
  dfi_config := workers.Configuration{"filename": "./test_data/dfi_test.txt"}
  dfi_node := &workers.DFINode{
    InputChannel: inputChannel,
    OutputChannel: outputChannel,
    ErrorChannel: errorChannel,
    ControlChannel: controlChannel,
    Config: dfi_config,
  }
  // init_doc := workers.Document{Value: []byte(""), Errors: []error{}}


  dfi_node.Setup()
  go dfi_node.Process()
  fmt.Println("waiting to finish processing...")
  time.Sleep(time.Second * 3)
  dfi_node.ControlChannel <- "exit"
  fmt.Println(<- dfi_node.OutputChannel)
  // dfo_node := &workers.DFONode{}
  // dfo_node.Input("./test_data/dfo_test.txt", []byte("dfo test here"), os.FileMode(int(0777)))
  // dfo_node.Process()
  // fmt.Println(string(dfi_node.Output().Value))
  fmt.Println("hello world")
}
