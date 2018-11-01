package workers

import (
	"fmt"
	"github.com/nsallis/elipse/log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// DFONode (Disk File Out) writes to a file from a given document value
// Takes config values:
//  - filepath - the path the document value should be written to
//  - append (optional) - if the file should be over-written/created, or appended to
//  - formatString (optional) - sets formatting for output file name where $SOURCE_NAME
// interpolates to the original file name

type DFONode struct {
	BaseNode
}

// GetNodeType get the type of node
func (n *DFONode) GetNodeType() string {
	return "DFO"
}

// ToString get the string version of this node
func (n *DFONode) ToString() string {
	return fmt.Sprintf("{UUID: %v, NodeType: %v, Config: %v, InputChannel: %v, OutputChannel: %v}", n.GetUUID(), n.GetNodeType(), n.GetConfig(), n.GetInput(), n.GetOutput())
}

// Setup make any config updates before processing
func (n *DFONode) Setup() {
	if str, ok := n.Config["filepath"]; ok {
		if !ok {
			log.Error("Required config `filepath` not present for node "+n.UUID, nil)
			panic("Required config `filepath` not preset for node")
		}
		absPath, err := filepath.Abs(str)
		if err != nil {
			log.Error("Could not get absolute path. Will try using relative path.", err)
			absPath = str // try using relative path
		}
		n.Config["filepath"] = absPath
	}
}

// Process run the worker. Writes the value of a document to the file. Can optionally
// append to the file (if append flag is true)
func (n *DFONode) Process() {
	var appendFlag bool
	if n.Config["append"] == "true" {
		appendFlag = true
	}

	for {
		select {
		case command := <-n.ControlChannel:
			if command == "exit" {
				log.Info("exiting node " + n.UUID)
				break
			}
		case document := <-n.InputChannel:
			filename := n.parseFileName(document.Source) // TODO will need to check if this is a sourceType other than disk
			fileInfo, err := os.Stat(filename)
			if err != nil {
				log.Error("Cannot find file for node "+n.UUID, err)
			}

			var file *os.File
			if appendFlag {
				file, err = os.OpenFile(filename, os.O_APPEND, fileInfo.Mode())
				if err != nil {
					log.Error("Cannot open file for node "+n.UUID, err)
				}
			} else {
				file, err = os.OpenFile(filename, os.O_CREATE, fileInfo.Mode())
				if err != nil {
					log.Error("Cannot open file for node "+n.UUID, err)
				}
			}
			_, err = file.Write(document.Value)
			file.Close()
			if err != nil {
				log.Error("Could not write to file for node "+n.UUID, err)
			}
		}
	}
}

func (n *DFONode) parseFileName(sourceName string) string { // TODO add more injectable values to this
	if formatString, ok := n.Config["formatString"]; ok {
		return path.Join(n.Config["filepath"], n.getFileNameFromPath(sourceName))
	} else {
		formattedFileName := strings.Replace(formatString, "$SOURCE_NAME", n.getFileNameFromPath(sourceName), -1)
		return path.Join(n.Config["filepath"], formattedFileName)
	}
}

// getFileNameFromPath gets the file name from file path (everything after last /)
func (n *DFONode) getFileNameFromPath(path string) string {
	pathSlice := strings.Split(path, "/")
	return pathSlice[len(pathSlice)-1]
}
