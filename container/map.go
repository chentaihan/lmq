package container

type Map map[int64]interface{}

func NewMap() *Map{
	return &Map{}
}

func (m *Map) Add(key int64, value interface{}) *Map{
	(*m)[key] = value
	return m
}

func (m Map) Get(key int64) interface{}{
	if val, ok := m[key]; ok {
		return val
	}
	return nil;
}

func (m Map) Contains(key int64) bool{
	_, ok := m[key]
	return ok
}

func (m *Map) Delete(key int64) *Map{
	delete(*m, key)
	return m
}

func (m *Map) Size() int{
	return len(*m)
}

func (m Map) Keys() []int64 {
	keys := make([]int64, m.Size());
	var index int = 0
	for key, _ := range m{
		keys[index] = key
		index++
	}
	return keys
}

func (m Map) Values() []interface{} {
	values := make([]interface{}, m.Size());
	var index int = 0
	for _, value := range m{
		values[index] = value
		index++
	}
	return values
}

