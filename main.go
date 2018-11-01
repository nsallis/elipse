package main

import (
	"fmt"

	"github.com/nsallis/elipse/spawner"
	// "github.com/nsallis/elipse/util"
	// "time"
	// "bufio"
	// "os"
)

func main() {

	config := spawner.CreateWorkerConfigFromFile("./test_data/worker_example.json")
	workersMap, _ := spawner.SpawnWorkers(config)
	spawner.ConnectWorkers(workersMap, config)
	for _, v := range workersMap {
		fmt.Println("starting node: " + v.ToString())
		go v.Process()
	}

	//***************************
	// inputChannel := make(chan workers.Document, 100)
	// outputChannel := make(chan workers.Document, 100)
	// controlChannel := make(chan string)
	// errorChannel := make(chan error, 100)
	// dfiConfig := workers.Configuration{"filename": "./test_data/dfi_test.txt"}
	// dfiNode := &workers.DFINode{
	// 	UUID:           "1",
	// 	InputChannel:   inputChannel,
	// 	OutputChannel:  outputChannel,
	// 	ErrorChannel:   errorChannel,
	// 	ControlChannel: controlChannel,
	// 	Config:         dfiConfig,
	// }

	// stdOutChan := make(chan workers.Document, 100)
	// stcCommandChan := make(chan string)
	// outNode := &workers.StdOutNode{
	// 	UUID:           "2",
	// 	InputChannel:   outputChannel,
	// 	OutputChannel:  stdOutChan,
	// 	ControlChannel: stcCommandChan,
	// 	ErrorChannel:   errorChannel,
	// 	Config:         workers.Configuration{},
	// }

	var command string

	// dfiNode.Setup()
	// outNode.Setup()
	// go dfiNode.Process()
	// go outNode.Process()
	fmt.Println("Started processing...")
	fmt.Println("")

	for { // main running loop
		fmt.Scanln(&command)
		if command != "" {
			// dfiNode.ControlChannel <- command
			// outNode.ControlChannel <- command
			if command == "exit" {
				fmt.Println("Waiting for nodes to exit...")
				break
			}
		}
	}
}
