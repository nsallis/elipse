package spawner

import (
	"encoding/json"
	"fmt"
	"github.com/nsallis/elipse/util"
	"github.com/nsallis/elipse/workers"
)

// CreateWorkerConfigFromFile creates a WorkerConfig from a file path
func CreateWorkerConfigFromFile(path string) []workers.WorkerConfig {
	jsonString, err := util.ReadDiskFile(path)
	if err != nil {
		panic(err)
	}
	return CreateWorkerConfig(jsonString)
}

// CreateWorkerConfig create worker config from json string
func CreateWorkerConfig(jsonString string) []workers.WorkerConfig { // needs to convert workers json into struct and return
	var config []workers.WorkerConfig
	if err := json.Unmarshal([]byte(jsonString), &config); err != nil {
		panic(err)
	}
	return config
}

// getNodeFromTypeString returns the type of node we need based on
// the type string. Also asserts the node type implements Node interface
func getNodeFromTypeString(typeName string) workers.Node {
	var node workers.Node
	switch typeName {
	case "DFI":
		node = &workers.DFINode{}
	case "StdOut":
		node = &workers.StdOutNode{}
	default:
		node = &workers.BaseNode{} // TODO this will eventually throw a not implemented
		// error because SetUUID is not implemented
	}
	return node
}

func spawnWorker(config workers.WorkerConfig) (workers.Node, error) {
	node := getNodeFromTypeString(config.NodeType)
	node.SetUUID(string(config.UUID))
	node.SetConfig(config.Config)
	return node, nil
}

// SpawnWorkers spawns workers from the config
func SpawnWorkers(configs []workers.WorkerConfig) (map[string]workers.Node, []error) {
	workersMap := make(map[string]workers.Node)
	for _, config := range configs {
		worker, _ := spawnWorker(config) // TODO log error
		workersMap[worker.GetUUID()] = worker
	}
	return workersMap, nil
}
