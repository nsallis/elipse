package spawner

import (
	"encoding/json"
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

// func getNodeFromTypeString(typeName string) workers.Node {
// 	switch typeName {
// 	case "DFI":
// 		return workers.DFINode{}
// 	}
// }
