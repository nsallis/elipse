package main

import (
	"fmt"

	"github.com/nsallis/elipse/log"
	"github.com/nsallis/elipse/server"
	"github.com/nsallis/elipse/spawner"
	// "github.com/nsallis/elipse/util"
	// "time"
	// "bufio"
	// "os"
)

func main() {

	log.InitLogs()
	log.Info("Reading worker configs")
	config := spawner.CreateWorkerConfigFromFile("./test_data/worker_example.json")
	workersMap, _ := spawner.SpawnWorkers(config)

	server := server.Server{}
	server.Initialize(workersMap)
	server.Start(3333)

	log.Info("Spawning workers")
	spawner.ConnectWorkers(workersMap, config)
	for _, v := range workersMap {
		log.Debug("starting node: " + v.ToString())
		v.Setup()
		go v.Process()
	}
	log.Info("finished spawning workers")

	var command string

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
