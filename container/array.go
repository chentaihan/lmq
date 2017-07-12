package container

import(
	"fmt"
)

type Array []interface{}

func NewArray() Array{
	return make(Array, 0, 4)
}

func NewArrayByCap(cap int) Array{
	return make(Array, 0, cap)
}

func (arr *Array) Append(val interface{}){
	*arr = append(*arr, val)
}

func (arr *Array) AppendRange(val []interface{}){
	*arr = append(*arr, val...)
}

func (arr *Array) Delete(index int) interface{}{
	if index < 0 || index >= arr.Size() {
		return nil
	}
	val := (*arr)[index]
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
	return val
}

func (arr *Array) DeleteRange(start, end int) bool{
	if start > end || start < 0 || end > arr.Size() {
		return false
	}
	*arr = append((*arr)[:start], (*arr)[end:]...)
	return true
}

func (arr Array) Get(index int) interface{}{
	return arr[index]
}

func (arr *Array) Set(index int, val interface{}) bool{
	if index >= 0 && index < len(*arr) {
		(*arr)[index] = val
		return true
	}
	return false
}

func (arr Array) Find(val interface{}) int{
	for index, item := range arr{
		if item == val {
			return index
		}
	}
	return -1
}

func (arr Array) Size() int{
	return len(arr)
}

func (arr Array) Cap() int{
	return cap(arr)
}

func ArrayTest(){
	arr1 := NewArrayByCap(0)
	fmt.Println(arr1.Size())
	fmt.Println(arr1.Cap())
	for i := 0; i < 10 ;i++  {
		arr1.Append(i)
	}
	arr1.Set(5,50)
	fmt.Println(arr1.Size())
	fmt.Println(arr1.Cap())
	slice1 := arr1[0:2]
	arr1.AppendRange(slice1)
	arr1.DeleteRange(0,arr1.Size()-5)
	for arr1.Size() > 0{
		fmt.Println(arr1.Get(0))
		arr1.Delete(0)
	}
}