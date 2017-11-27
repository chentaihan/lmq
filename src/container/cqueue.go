package container

import "sync"

type CQueue struct {
	list []interface{}
	rwLock sync.RWMutex
}

func NewCQueue() *CQueue{
	return &CQueue{list:make([]interface{},0)}
}

func (s *CQueue) Enqueue(item interface{}) {
	s.rwLock.Lock()
	s.list = append(s.list, item)
	s.rwLock.Unlock()
}

func (s *CQueue) Dequeue() interface{}{
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if len(s.list) > 0 {
		ret := s.list[0]
		s.list = s.list[1:]
		return ret
	}
	return nil
}

func (s *CQueue) Size() int{
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return len(s.list)
}

func (s *CQueue) Peek() interface{}{
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	if len(s.list) > 0 {
		return s.list[0]
	}
	return nil
}

func (s *CQueue) Copy() []interface{}{
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	if len(s.list) > 0 {
		list := s.list[0:]
		s.list = make([]interface{}, 0)
		return list
	}
	return nil
}

