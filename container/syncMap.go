package container

import (
	"fmt"
	"sync"
)

type AsyncMap map[int64]interface{}

type SyncMap struct {
	Map AsyncMap
	rwLock sync.RWMutex
}

func NewSyncMap() *SyncMap{
	return &SyncMap{Map : make(AsyncMap)}
}

func (m *SyncMap) Add(key int64, value interface{}) *SyncMap{
	m.rwLock.Lock();
	m.Map[key] = value
	m.rwLock.Unlock();
	return m
}

func (m SyncMap) Get(key int64) interface{}{
	m.rwLock.RLock();
	val, ok := m.Map[key];
	m.rwLock.RUnlock();
	if ok {
		return val
	}
	return nil;
}

func (m SyncMap) Contains(key int64) bool{
	m.rwLock.RLock()
	_, ok := m.Map[key]
	m.rwLock.RUnlock()
	return ok
}

func (m SyncMap) ContainsValue(value interface{}) bool{
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	for _, val := range m.Map{
		if value == val{
			return true
		}
	}
	return false
}

func (m *SyncMap) Delete(key int64) *SyncMap{
	m.rwLock.Lock();
	delete(m.Map, key)
	m.rwLock.Unlock()
	return m
}

func (m *SyncMap) Size() int{
	m.rwLock.RLock();
	size:= len(m.Map)
	m.rwLock.RUnlock();
	return size
}

func (m SyncMap) Keys() []int64 {
	keys := make([]int64, m.Size());
	var index int = 0
	m.rwLock.RLock();
	for key, _ := range m.Map{
		keys[index] = key
		index++
	}
	m.rwLock.RUnlock()
	return keys
}

func (m SyncMap) Values() []interface{} {
	values := make([]interface{}, m.Size());
	var index int = 0
	m.rwLock.RLock();
	for _, value := range m.Map{
		values[index] = value
		index++
	}
	m.rwLock.RUnlock()
	return values
}

func SyncMapTest(){
	m := NewSyncMap()
	fmt.Println(m.Size())
	m.Add(1, "value value111111")
	m.Add(2, "ffffffffff11111111")
	m.Add(3, "int64")

	fmt.Println(m.Get(1))
	fmt.Println(m.Get(2))
	fmt.Println(m.Size())
	fmt.Println(m.Contains(3))
	fmt.Println(m.Contains(4))
	fmt.Println(m.Delete(3).Size());
 	fmt.Println(m.ContainsValue("int64"));
	for  index, key := range m.Keys(){
		fmt.Print(index)
		fmt.Print("    ")
		fmt.Println(key)
	}
	for index, val := range m.Values(){
		fmt.Print(index)
		fmt.Print("    ")
		fmt.Println(val)
	}
	fmt.Println("key=>value");
	for key, value := range m.Map{
		fmt.Print(key)
		fmt.Print("    ")
		fmt.Println(value)
	}
}
