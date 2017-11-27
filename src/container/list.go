package container

import (
	"fmt"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
	len  int
}

func NewList() *List {
	return new(List)
}

func (list List) Front() interface{} {
	if list.tail != nil {
		return list.tail.Value;
	}
	return nil
}

func (list List) Back() interface{}{
	if list.tail != nil {
		return list.tail.Value;
	}
	return nil
}

func (list List) Begin() *ListNode {
	return list.head
}

func (list List) End() *ListNode{
	return list.tail
}

func (list *List) PushFront(val interface{}) *ListNode {
	node := new(ListNode)
	node.Value = val
	if list.head == nil {
		list.head, list.tail = node, node
	} else {
		node.Next = list.head;
		list.head = node;
	}
	list.len++
	return node
}

func (list *List) PushBack(val interface{}) *ListNode {
	node := new(ListNode)
	node.Value = val
	if list.head == nil {
		list.head, list.tail = node, node
	} else {
		list.tail.Next = node
		list.tail = node
	}
	list.len++
	return node
}


func (list *List) Insert(cur *ListNode, val interface{}) *ListNode {
	if cur != nil {
		node := new(ListNode)
		node.Value = val
		node.Next = cur.Next
		cur.Next = node
		list.len++
		return cur
	} else {
		return list.PushBack(val)
	}
}

func (list *List) InsertIndex(index int, val interface{}) *ListNode {
	curNode := list.head
	if list.len > index {
		for i := 0; i < index; i++ {
			curNode = curNode.Next
		}
		return list.Insert(curNode, val)
	}
	return list.PushBack(val)
}

func (list *List) RemoveValue(val interface{}) {
	if list.head == nil {
		return
	}
	index := list.Index(val)
	if index >= 0{
		list.RemoveIndex(index)
	}
}

func (list *List) RemoveIndex(index int) *ListNode {
	if list.len == 0 || index < 0 || index >= list.len {
		return nil
	}

	if index == 0{
		tmpNode := list.head
		list.head = list.head.Next
		if list.len == 1 {
			list.tail = list.head
		}
		list.len--
		return tmpNode
	}

	prevNode := list.IndexValue(index - 1)
	curNode := prevNode.Next
	if curNode == nil {
		return nil
	}
	prevNode.Next = curNode.Next
	if prevNode.Next == nil {
		list.tail = prevNode
	}
	list.len--
	return curNode
}

func (list *List) Clear()  {
	list.head, list.tail = nil, nil
	list.len = 0
}

func (list List) Index(val interface{}) int {
	index := 0
	for curNode := list.head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == val {
			return index
		}
		index++
	}
	return -1
}

func (list List) IndexValue(index int) *ListNode {
	if index < 0 || index >= list.len {
		return nil
	}
	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}
	return curNode
}

func (list List) Size() int {
	return list.len
}

func (list *List) Print() {
	tmpList := make([]interface{}, list.len)
	index := 0
	for curNode := list.Begin(); curNode != nil; curNode = curNode.Next {
		tmpList[index] = curNode.Value
		index++
	}
	fmt.Println(tmpList)
}

func Print(list *ListNode) {
	tmpArr := make([]interface{}, 0)
	for ; list != nil; list = list.Next {
		tmpArr = append(tmpArr, list.Value)
	}
	fmt.Println(tmpArr)
}

func ListTest() {
	list := NewList()
	fmt.Println(list.Front())
	fmt.Println(list.Begin())
	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}
	list.PushFront(0)
	list.Print()
	fmt.Println(list.RemoveIndex(6))
	fmt.Println(list.IndexValue(3))
	list.RemoveValue(5)
	list.Print()
	fmt.Println(list.Index(5))
	fmt.Println(list.Index(4))
	list.InsertIndex(50, 51)

	list.Print()
}
