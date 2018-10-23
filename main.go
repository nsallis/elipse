package main

import (
  "fmt"
  "github.com/nsallis/elipse/workers"
  // "time"
  // "bufio"
  // "os"
)

func main() {
  inputChannel := make(chan workers.Document, 100)
  outputChannel := make(chan workers.Document, 100)
  controlChannel := make(chan string)
  errorChannel := make(chan error, 100)
  dfi_config := workers.Configuration{"filename": "./test_data/dfi_test.txt"}
  dfi_node := &workers.DFINode{
    UUID: "1",
    InputChannel: inputChannel,
    OutputChannel: outputChannel,
    ErrorChannel: errorChannel,
    ControlChannel: controlChannel,
    Config: dfi_config,
  }

  stdOutChan := make(chan workers.Document, 100)
  stcCommandChan := make(chan string)
  out_node := &workers.StdOutNode{
    UUID: "2",
    InputChannel: outputChannel,
    OutputChannel: stdOutChan,
    ControlChannel: stcCommandChan,
    ErrorChannel: errorChannel,
    Config: workers.Configuration{},
  }

  var command string

  dfi_node.Setup()
  out_node.Setup()
  go dfi_node.Process()
  go out_node.Process()
  fmt.Println("Started processing...\n")

  for { // main running loop
    fmt.Scanln(&command)
    if(command != "") {
      dfi_node.ControlChannel <- command
      out_node.ControlChannel <- command
      if(command == "exit") {
        fmt.Println("Waiting for nodes to exit...")
        break
      }
    }
  }
}
