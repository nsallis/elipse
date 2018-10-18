package workers

// disk file output

import (
  "io/ioutil"
  "path/filepath"
  "os"
)

type DFONode struct {
  filename string
  value []byte
  permission os.FileMode
  err error
}

func (n *DFONode) Input(filename string, value []byte, permission os.FileMode) { // will eventually take a document
  // var _ Node = DFINode{}
  n.filename = filename
  n.value = value
  n.permission = permission
}

func (n *DFONode) Process() {
  filename, _ := filepath.Abs(n.filename)

  err := ioutil.WriteFile(filename, n.value, n.permission)
  n.err = err
}

func (n *DFONode) Output() []byte { // will eventually return document
  return n.value
}

func (n *DFONode) Error() NodeError { // if there is an error operating on this node, call this
  return n.err
}
