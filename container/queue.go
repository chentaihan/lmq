package container

import (
"fmt"
)

type Queue struct{
	list List
}

func NewQueue() *Queue{
	return new(Queue)
}

func (s *Queue) Enqueue(item interface{}) {
	s.list.PushBack(item)
}

func (s *Queue) Dequeue() interface{}{
	return s.list.RemoveIndex(0).Value
}

func (s *Queue) Size() int{
	return s.list.Size()
}

func (s *Queue) Peek() interface{}{
	return s.list.IndexValue(0).Value
}

func (s *Queue) Clear() {
	s.list.Clear()
}

func (s *Queue) Contains(item interface{}) bool{
	return s.list.Index(item) >= 0
}

func QueueTest() {
	st := NewQueue()
	for i := 0; i < 10; i++ {
		st.Enqueue(i)
	}
	fmt.Println(st.Size())
	st.Clear()
	for i := 0; i < 10; i++ {
		st.Enqueue(i)
	}
	fmt.Println(st.Size())
	fmt.Println(st.Contains(5))
	for st.Size() > 0 {
		fmt.Print(st.Dequeue(), " ")
	}
	fmt.Println("over")
}