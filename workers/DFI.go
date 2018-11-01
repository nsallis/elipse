package workers

// disk file input

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nsallis/elipse/log"
	"io/ioutil"
	"path/filepath"
	"time"
)

// DFINode reads files from disk, and passes them on as documents
type DFINode struct {
	BaseNode
}

// GetNodeType get the type of node
func (n *DFINode) GetNodeType() string {
	return "DFI"
}

// ToString get the string version of this node
func (n *DFINode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}

// Setup make any config updates before processing
func (n *DFINode) Setup() {
	// var _ Node = DFINode{}
	// TODO check if file path is valid
	// TODO get the absolute path of the file
	if str, ok := n.Config["filename"]; ok {
		absPath, err := filepath.Abs(str)
		if err != nil {
			log.Error("Could not get absolute path. Will try using relative path.", err)
			absPath = str // try using relative path
		}
		n.Config["filename"] = absPath
	}
}

// Process run the worker. Reads the specified file in and converts it to a document.
// Also watches the specified file for any changes and outputs an updated document if there have been changes
func (n *DFINode) Process() {
	var filepath string
	if str, ok := n.Config["filename"]; ok {
		filepath = str
	}

	if str, ok := n.Config["filename"]; ok {
		fileContents, _ := ioutil.ReadFile(str)
		n.OutputChannel <- n.createDocument(n, fileContents)
	}

	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	watcher.Add(filepath)
	log.Info("node " + n.UUID + " is watching file/directory " + filepath)
	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				break
			}
		case event := <-watcher.Events:
			switch {
			case event.Op == fsnotify.Write:
				fmt.Println("wrote to the file")
				fileContents, _ := ioutil.ReadFile(event.Name)
				n.OutputChannel <- n.createDocument(n, fileContents)
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

func (n *DFINode) createDocument(node *DFINode, fileContents []byte) Document {

	var filepath string
	if str, ok := node.Config["filename"]; ok {
		filepath = str
	}
	return Document{Value: fileContents, Source: filepath, SourceType: "disk"}
}
