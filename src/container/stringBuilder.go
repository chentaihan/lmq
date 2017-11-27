package container

//尝试将切片封装成StringBuilder，结果速度更慢，那还封装个毛线

import (
	"fmt"
	"time"
)

type StringBuilder struct {
	Buf []byte
	Cap int
}

func NewStringBuilder(cap int) *StringBuilder{
	if cap < 64 {
		cap = 64;
	}
	sb := new(StringBuilder)
	sb.Cap = cap
	sb.Buf = make([]byte, 0, cap)
	return sb
}

func (sb *StringBuilder) AppendBytes(str []byte){
	tmpSize := len(str) + len(sb.Buf)
	if sb.Cap < tmpSize {
		sb.Cap = tmpSize*2
		newBuf := make([]byte, len(sb.Buf), sb.Cap)
		copy(newBuf, sb.Buf)
		sb.Buf = newBuf
	}
	sb.Buf = append(sb.Buf, str...)
}

func (sb *StringBuilder) ToString() string{
	return string(sb.Buf)
}

func (sb *StringBuilder) GetBytes() []byte{
	return sb.Buf
}

func (sb *StringBuilder) AppendByte(c byte){
	sb.Buf = append(sb.Buf, c)
}

func TestStringBuilder(){
	count := 10000
	startTime1 := time.Now().UnixNano()
	byte1 := []byte("123456789qwertyuiopasdfghjkl;'zxcvbnm,./")
	byte2 := []byte("___123456789qwertyuiopasdfghjkl;'zxcvbnm,./陈太汉")
	for i := 0; i < count; i++ {
		sb := NewStringBuilder(256)
		sb.AppendByte('[')
		for i := 0; i < 10; i++ {
			sb.AppendBytes(byte1)
			sb.AppendBytes(byte2)
		}
		sb.AppendByte(']')
	}
	endTime1 := time.Now().UnixNano() - startTime1
	fmt.Println(endTime1)

	startTime1 = time.Now().UnixNano()
	for i := 0; i < count; i++ {
		sb1 := make([]byte, 0, 64)
		sb1 = append(sb1, '[')
		for i := 0; i < 10; i++ {
			sb1 = append(sb1, byte1...)
			sb1 = append(sb1, byte2...)
		}
		sb1 = append(sb1, ']')
	}
	endTime1 = time.Now().UnixNano() - startTime1
	fmt.Println(endTime1)

	sb := NewStringBuilder(256)
	sb.AppendByte('[')
	for i := 0; i < 10; i++ {
		sb.AppendBytes(byte1)
		sb.AppendBytes(byte2)
	}
	sb.AppendByte(']')
	fmt.Println(sb.ToString())
	fmt.Println(sb.Cap)
}