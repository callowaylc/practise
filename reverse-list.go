package main

import (
  _ "fmt"
  "log"
  _ "io/ioutil"
  "os"
)


// type /////////////////////////////////////////

type Node struct{
  Next *Node
  Value string
}

// methods //////////////////////////////////////

func (n Node) String() string {
  return n.Value
}

func (n *Node) Add(value string) *Node {
  n.Next = &Node{
    Value: value,
  }
  return n.Next
}

// variables ////////////////////////////////////

var logger *log.Logger

// functions ////////////////////////////////////

func init() {
  logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
  head := Node{
    Value: "one",
  }
  head.Add("two").Add("three")

  printLinkedList(reverseLinkedList(&head, nil))

}

func reverseLinkedList(list *Node, previous *Node) *Node  {
  current_head := list.Next
  list.Next = previous

  logger.Printf("current_head: %+v", current_head)
  logger.Printf("list: %v", list)
  logger.Println("---")

  if current_head != nil {
    return reverseLinkedList(current_head, list)
  }

  return list
}


func printLinkedList(head *Node) {
  node := head
  hasNext := node.Next != nil

  for hasNext {
    logger.Println(node)
    hasNext = node.Next != nil
    node = node.Next
  }
}
