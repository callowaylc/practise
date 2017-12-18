package main

import (
  "fmt"
)

// interface ////////////////////////////////////

type LinkedListNode interface {
  HasNext() bool
  Add(value string) *Node
}

// type /////////////////////////////////////////

type Node struct{
  Next *Node
  Value string
}

// methods //////////////////////////////////////

func (n *Node) HasNext() bool {
  return n != nil
}

func (n *Node) Add(value string) *Node {
  next := &Node{
    Value: value,
  }
  n.Next = next

  return next
}

// functions ////////////////////////////////////

func init() { }

func main() {
  head := Node{
    Value: "f",
  }
  head.Add("u").Add("n")

  printLinkedList(&head)
}

//func reverseLinkedList(head LinkedListNode) LinkedListNode {
//
//}

func printLinkedList(head Node) {
  node := head.(Node)
  for node.HasNext() {
    fmt.Printf("%s", node)
    node = node.Next
  }
}
