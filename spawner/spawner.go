package spawner

import (
	"encoding/json"
	"fmt"
	"github.com/nsallis/elipse/log"
	"github.com/nsallis/elipse/util"
	"github.com/nsallis/elipse/workers"
	"time"
)

// WorkerConfig configuration for all workers.
// Used in building nodes from json
// TODO should probably be moved to spawner
type WorkerConfig struct {
	UUID     string
	NodeType string
	Config   map[string]string
	Outputs  []string
	Errors   []string
}

// CreateWorkerConfigFromFile creates a WorkerConfig from a file path
func CreateWorkerConfigFromFile(path string) []WorkerConfig {
	jsonString, err := util.ReadDiskFile(path)
	if err != nil {
		panic(err)
	}
	return CreateWorkerConfig(jsonString)
}

// CreateWorkerConfig create worker config from json string
func CreateWorkerConfig(jsonString string) []WorkerConfig { // needs to convert workers json into struct and return
	var config []WorkerConfig
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
	case "DFO":
		node = &workers.DFONode{}
	case "Splitter":
		node = &workers.SplitterNode{}
	case "GoProcessor":
		node = &workers.GoProcessorNode{}
	case "Joiner":
		node = &workers.JoinerNode{}
	default:
		node = &workers.BaseNode{} // TODO this will eventually throw a not implemented
		// error because SetUUID is not implemented
	}
	return node
}

func spawnWorker(config WorkerConfig) (workers.Node, error) {
	node := getNodeFromTypeString(config.NodeType)
	node.SetUUID(string(config.UUID))
	node.SetConfig(config.Config)
	return node, nil
}

// SpawnWorkers spawns workers from the config
func SpawnWorkers(configs []WorkerConfig) (map[string]workers.Node, []error) {
	workersMap := make(map[string]workers.Node)
	for _, config := range configs {
		worker, _ := spawnWorker(config) // TODO log error
		workersMap[worker.GetUUID()] = worker
	}
	return workersMap, nil
}

func ConnectWorkers(workersMap map[string]workers.Node, configs []WorkerConfig) {
	for _, v := range workersMap {
		inChannel := make(chan workers.Document, 100) // TODO make the channel buffer length configurable
		outChannel := make(chan workers.Document, 100)
		contChannel := make(chan string)
		errChannel := make(chan error)
		v.SetInput(inChannel)
		v.SetOutput(outChannel)
		v.SetControl(contChannel)
		v.SetError(errChannel)
	}

	var i int
	// TODO instead, loop through config, and pull the workers by UUID
	for i := 0; i < len(configs); i++ {
		currentConfig := configs[i]
		currentWorker := workersMap[currentConfig.UUID]
	}
	for _, v := range workersMap {
		outputUUIDS := configs[i].Outputs
		fmt.Println("currentWorker: %v", v)
		if len(outputUUIDS) > 0 {
			fmt.Println("corresponding worker: %v", workersMap[outputUUIDS[0]])
		}
		log.Debug("")
		// log.Debug(fmt.Sprintf("i: %v", i))
		// log.Debug(fmt.Sprintf("current node: %v", v))
		// time.Sleep(3 * time.Second)
		// log.Debug(fmt.Sprintf("hooking up worker: %i, with worker %i", i, outputUUIDS[0]))
		// log.Debug(fmt.Sprintf("pointing %v to %v", v, workersMap[outputUUIDS[0]]))
		if len(outputUUIDS) > 0 {
			attachedNode := workersMap[outputUUIDS[0]] // just use the first output for now
			v.SetOutput(attachedNode.GetInput())
		}
		i++
	}
	time.Sleep(3 * time.Second)
}
