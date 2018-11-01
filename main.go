package main

import (
	"fmt"

	"errors"
	"github.com/nsallis/elipse/log"
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
	fmt.Println("loggin an error for testing: ")
	log.Error("This is a test error", errors.New("this is a test stack trace"))

	var command string

	fmt.Println("Started processing...")
	fmt.Println("")

	for { // main running loop
		fmt.Scanln(&command)
		if command != "" {
			for _, v := range workersMap {
				v.GetControl() <- command
			}
			if command == "exit" {
				fmt.Println("Waiting for nodes to exit...")
				break
			}
		}
	}
}
