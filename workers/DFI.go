package workers

// disk file input

import (
  "io/ioutil"
  "path/filepath"
)

type DFINode struct {
  filename string
  value string
  err error
}

func (n *DFINode) Input(d Document, config Configuration) { // will eventually take a document
  // var _ Node = DFINode{}
  // TODO get file name from configuration
  if str, ok := config["filename"].(string); ok {
    n.filename = str
  } else {
    // TODO throw bad config
    /* not string */
  }
}

func (n *DFINode) Process() {
  filename, _ := filepath.Abs(n.filename)

  fileContents, err := ioutil.ReadFile(filename) // TODO can't find the file for some reason!
  n.value = string(fileContents)
  n.err = err
}

func (n *DFINode) Output() Document { // will eventually return document
  returnDoc := Document{Value: []byte(n.value)}
  return returnDoc
}

func (n *DFINode) Error() NodeError { // if there is an error operating on this node, call this
  return n.err
}
