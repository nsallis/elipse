// STD out logger. Logs input documents to stdout
package workers

import (
  "fmt"
  "time"
)

type StdOutNode NodeStruct

func (n *StdOutNode) Setup() {

}

func (n *StdOutNode) Process() {
  for {
    select {
    case inputDoc := <- n.InputChannel:
      docString, _ := inputDoc.ToString()
      fmt.Println("Got a document: " + docString)
    case command := <- n.ControlChannel:
      if(command == "exit"){
        fmt.Println("exiting stdout...")
        close(n.OutputChannel)
        break
      }

    default:
      time.Sleep(time.Millisecond * 100)
    }
  }
}
