package main

import (
	"fmt"
	"errors"
)

type List struct {
	length int
	firstNode *Node
	lastNode *Node
}

type Node struct {
	pos uint64
	data int
	next *Node
}

func main(){
	list := &List{}
	
	n1 := Node{
		pos: 0,
		data: 1996,
	}
	n2 := Node{
		pos: 1,
		data: 1997,
	}
	list.Append(&n1)
	list.Append(&n2)
	
	fmt.Printf("Length: %v\n", list.length)
	fmt.Printf("First: %v\n", list.firstNode)
	fmt.Printf("Second: %v\n", list.firstNode.next)
	
	n3 := Node{
		pos:3,
		data: 1999,
	}
	//append at the end of the list (pos not taken into account)
	list.Append(&n3)
	
	fmt.Printf("Length: %v\n", list.length)
	fmt.Printf("Third: %v\n", list.firstNode.next.next)
	
	n4 := Node{
		pos:2,
		data: 1998,
	}
	//insert between pos 1 and 3
	list.Insert(&n4)

	fmt.Printf("Length: %v\n", list.length)
	fmt.Printf("Third: %v\n", list.firstNode.next.next)

	//loop through list and print each node's data
	node := list.firstNode
	for node != nil {
		fmt.Printf("%v\n", node.data)
		node = node.next
	}
}

//returns head of list (first node)
func (list *List) Head() (*Node, bool){
	return list.firstNode, true
}

//appends new node to end of the list
func (list *List) Append(newNode *Node) {
	if list.length == 0 {
		list.firstNode = newNode
		list.lastNode = newNode
	} else {
		currentNode := list.lastNode
		currentNode.next = newNode
		list.lastNode = newNode
	}
	list.length++
}

//removes node from list at given position
func (list *List) Remove(pos uint64) {
	if list.length == 0 {
		panic(errors.New("List is empty"))
	}

	var previousNode *Node
	currentNode := list.firstNode

	for currentNode.pos != pos {
		if currentNode.next == nil {
			panic(errors.New("Node doesn't exist"))
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}
	previousNode.next = currentNode.next

	list.length--
}

//insert new node into list at correct position
func (list *List) Insert(newNode *Node){
	if list.length == 0 {
		list.firstNode = newNode
	} else {
		var previousNode *Node 
		currentNode := list.firstNode

		for currentNode.pos < newNode.pos {
			previousNode = currentNode
			currentNode = previousNode.next
			
		}

		previousNode.next = newNode
		newNode.next = currentNode
	}
	list.length++
}
