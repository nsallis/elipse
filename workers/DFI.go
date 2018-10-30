package workers

// disk file input

import (
	"io/ioutil"
	// // "path/filepath"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"time"
)

type DFINode struct {
	BaseNode
}

func (n DFINode) GetNodeType() string {
	return "DFI"
}

func (n DFINode) SetUUID(uuid string) {
	n.UUID = uuid
}

func (n DFINode) SetConfig(config map[string]string) {
	n.Config = config
}

func (n DFINode) Setup() {
	// var _ Node = DFINode{}
	// TODO check if file path is valid
	// TODO get the absolute path of the file
	if str, ok := n.Config["filename"]; ok {
		n.Config["filename"] = str
	}
}

func (n DFINode) Process() {
	var filepath string
	if str, ok := n.Config["filename"]; ok {
		filepath = str
	}

	if str, ok := n.Config["filename"]; ok {
		fileContents, _ := ioutil.ReadFile(str)
		n.OutputChannel <- createDocFromNode(n, fileContents)
	}

	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	watcher.Add(filepath)
	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				fmt.Println("exiting DFI...")
				break
			}
		case event := <-watcher.Events:
			switch {
			case event.Op == fsnotify.Write:
				fmt.Println("wrote to the file")
				fileContents, _ := ioutil.ReadFile(event.Name)
				n.OutputChannel <- createDocFromNode(n, fileContents)
			case event.Op == fsnotify.Create:
				fmt.Println("created a file in a watched directory")
			default:

			}
			fmt.Println(event)
		default:
			time.Sleep(time.Millisecond * 100) // dictates responsiveness
		}
	}
}

func createDocFromNode(node DFINode, fileContents []byte) Document {

	var filepath string
	if str, ok := node.Config["filename"]; ok {
		filepath = str
	}
	return Document{Value: fileContents, Source: filepath, SourceType: "disk"}
}
