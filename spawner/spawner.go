package spawner

import (
	"encoding/json"
	"github.com/nsallis/elipse/util"
)

// WorkerConfig configuration for all workers
type WorkerConfig interface{}

// CreateWorkerConfigFromFile creates a WorkerConfig from a file path
func CreateWorkerConfigFromFile(path string) WorkerConfig {
	jsonString, err := util.ReadDiskFile(path)
	if err != nil {
		panic(err)
	}
	return CreateWorkerConfig(jsonString)
}

// CreateWorkerConfig create worker config from json string
func CreateWorkerConfig(jsonString string) WorkerConfig { // needs to convert workers json into struct and return
	var config WorkerConfig
	if err := json.Unmarshal([]byte(jsonString), &config); err != nil {
		panic(err)
	}
	return config
}
