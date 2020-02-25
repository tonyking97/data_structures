package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type List struct {
	len   int
	start *node
}

func (l *List) addNode(data int) {
	newNode := &node{
		data: data,
		next: nil,
	}
	if l.len == 0 {
		//root node
		l.start = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}

	l.len++
}

func (l *List) addRootNode(data int) {
	newNode := &node{
		data: data,
		next: l.start,
	}
	l.start = newNode
}

func (l *List) addNodeArray(data []int) {
	for _, num := range data {
		l.addNode(num)
	}
}

func (l *List) insertNode(pos, data int) {
	if pos > l.len || pos < 0 {
		return
	} else {
		currentNode := l.start
		previousNode := &node{}
		position := 0
		for currentNode != nil {
			if pos == position {
				newNode := &node{
					data: data,
					next: currentNode,
				}
				previousNode.next = newNode
				break
			}
			position++
			previousNode = currentNode
			currentNode = currentNode.next
		}
	}
}

func (l *List) length() int {
	return l.len
}

func (l *List) isEmpty() bool {
	return l.len == 0
}

func (l *List) find(data int) int {
	currentNode := l.start
	position := 0
	for currentNode != nil {
		if currentNode.data == data {
			return position
		}
		position++
		currentNode = currentNode.next
	}
	return -1
}

func (l *List) contains(data int) bool {
	currentNode := l.start
	for currentNode != nil {
		if currentNode.data == data {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (l *List) updateNode(pos, data int) {
	if pos >= l.len || pos < 0 {
		return
	} else {
		currentNode := l.start
		position := 0
		for currentNode != nil {
			if position == pos {
				currentNode.data = data
				break
			}
			position++
			currentNode = currentNode.next
		}
	}
}

func (l *List) deleteNode(pos int) {
	if pos >= l.len || pos < 0 {
		return
	} else if pos == 0 {
		l.start = l.start.next
	} else {
		currentNode := l.start
		previousNode := &node{}

		position := 0
		for currentNode != nil {
			if pos == position {
				previousNode.next = currentNode.next
				currentNode = nil //deleting the variable to free memory {Taken care of GC}
				break
			}
			position++
			previousNode = currentNode
			currentNode = currentNode.next
		}
	}
	l.len--
}

func (l *List) reverse() {
	currentNode := l.start
	previousNode := &node{}
	previousNode = nil

	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	l.start = previousNode
}

func (l *List) display() {
	fmt.Println("---------------")
	currentNode := l.start
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("---------------")
}

func main() {
	list := &List{}
	//list.addNode(1)
	//list.addNode(2)
	//list.addNode(3)
	//list.addNode(4)
	//list.addNode(5)
	//list.addNode(6)
	//list.display()
	//fmt.Println("Length -> ",list.length())
	//list.deleteNode(0)
	//list.deleteNode(3)
	//list.deleteNode(5)
	//list.display()
	//fmt.Println("Length -> ",list.length())
	nodeArr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list.addNodeArray(nodeArr)
	//list.display()
	fmt.Println("Length -> ", list.length())
	list.updateNode(5, 9)
	list.addRootNode(0)
	list.display()
	fmt.Println(list.contains(10))
	list.reverse()
	list.display()
}
