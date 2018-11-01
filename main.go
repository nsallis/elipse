package main

import (
	"fmt"

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

	var command string

	log.Info("Starting processing...")

	for { // main running loop
		fmt.Scanln(&command)
		if command != "" {
			for _, v := range workersMap {
				v.GetControl() <- command
			}
			if command == "exit" {
				break
			}
		}
	}
	log.Info("All nodes have stopped. Shutting down.")
}
