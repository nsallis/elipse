package main

import (
  "fmt"
  "github.com/nsallis/elipse/workers"
  "os"
)

func main() {
  dfi_node := &workers.DFINode{}
  init_doc := workers.Document{Value: []byte(""), Errors: []error{}}
  dfi_config := workers.Configuration{"filename": "./test_data/dfi_test.txt"}
  dfi_node.Input(init_doc, dfi_config)
  dfi_node.Process()
  dfo_node := &workers.DFONode{}
  dfo_node.Input("./test_data/dfo_test.txt", []byte("dfo test here"), os.FileMode(int(0777)))
  dfo_node.Process()
  fmt.Println(string(dfi_node.Output().Value))
  fmt.Println("hello world")
}
