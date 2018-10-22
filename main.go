package main

import (
  "fmt"
  "github.com/nsallis/elipse/workers"
  // "time"
  // "bufio"
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

  var command string


  dfi_node.Setup()
  go dfi_node.Process()
  fmt.Println("waiting to finish processing...")

  for { // main running loop
    fmt.Println( <- outputChannel ) // we need to pull from the channel so we don't block
    fmt.Scanln(&command)
    if(command != "") {
      if(command == "exit") {
        dfi_node.ControlChannel <- command
        break

      }
    }
  }
  // fmt.Println(<- dfi_node.OutputChannel)
  // dfo_node := &workers.DFONode{}
  // dfo_node.Input("./test_data/dfo_test.txt", []byte("dfo test here"), os.FileMode(int(0777)))
  // dfo_node.Process()
  // fmt.Println(string(dfi_node.Output().Value))
}
