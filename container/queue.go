package container

type Queue interface {
	Enqueue(item interface{})
	Dequeue() interface{}
	Size() int
	Peek() interface{}
	Clear()
}

