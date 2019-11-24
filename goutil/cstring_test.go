package goutil

import (
	"fmt"
	"testing"
)

func TestGetGStrFromCStr(t *testing.T) {
	data := []byte{0x68, 0x65, 0x6c, 0x00, 0x6c, 0x6f}
	strVal := GetGStrFromCStr(data)
	fmt.Println(strVal)

	data2 := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	strVal2 := GetGStrFromCStr(data2)
	fmt.Println(strVal2)

	data3 := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00}
	strVal3 := GetGStrFromCStr(data3)
	fmt.Println(strVal3)
}

func TestGetCStrFromGStr(t *testing.T) {
	uuid := "hello"
	str := GetCStrFromGStr(uuid)
	fmt.Println(str)
}
