package math

import (
	"bytes"
	"encoding/binary"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BytesToInt(byteArr []byte) int {
	oldLen := len(byteArr)
	if oldLen < 8 {
		newArr := make([]byte, 8-oldLen)
		byteArr = append(newArr, byteArr...)
	} else if oldLen > 8 {
		byteArr = byteArr[oldLen-8:]
	}
	var i int64
	binary.Read(bytes.NewBuffer(byteArr), binary.BigEndian, &i)
	return int(i)
}

func IntToBytes(i int, byteLen int) (byteArr []byte) {
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, int64(i))
	byteArr = buffer.Bytes()
	if byteLen < 8 {
		byteArr = byteArr[8-byteLen:]
	} else if byteLen > 8 {
		newArr := make([]byte, byteLen-8)
		byteArr = append(newArr, byteArr...)
	}
	return byteArr
}
