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
	a := 0
	binary.Read(bytes.NewBuffer(byteArr), binary.BigEndian, &a)
	return a
}
