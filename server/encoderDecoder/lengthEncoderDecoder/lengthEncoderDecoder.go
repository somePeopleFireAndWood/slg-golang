package lengthEncoderDecoder

import (
	"slg-golang/math"
)

type LengthDecoder struct {
	ReadState       int // 1 正在读取长度头； 2 正在读取消息本体
	LastRemainBytes []byte
	ReadLen         int // 用来控制消息体长度的head长度
}

var UseLengthHeadLen = 4

func NewLengthDecoder() *LengthDecoder {
	decoder := LengthDecoder{
		ReadState:       1,
		ReadLen:         UseLengthHeadLen,
		LastRemainBytes: make([]byte, 0, UseLengthHeadLen),
	}
	return &decoder
}

func (d *LengthDecoder) Decode(in []byte) (outArr [][]byte) {
	outArr = [][]byte{}
	readIndex := 0
	inLen := len(in)
	for readIndex < inLen {
		if d.ReadState == 1 {
			d.tryReadBytes(in, &readIndex)
			if len(d.LastRemainBytes) == d.ReadLen {
				d.ReadLen = math.BytesToInt(d.LastRemainBytes)
				d.ReadState = 2
				d.LastRemainBytes = make([]byte, 0, d.ReadLen)
			}
		} else if d.ReadState == 2 {
			d.tryReadBytes(in, &readIndex)
			if len(d.LastRemainBytes) == d.ReadLen {
				outArr = append(outArr, d.LastRemainBytes)
				d.ReadLen = UseLengthHeadLen
				d.ReadState = 1
				d.LastRemainBytes = make([]byte, 0, d.ReadLen)
			}
		}
	}
	return outArr
}

func (d *LengthDecoder) tryReadBytes(in []byte, readIndex *int) {
	inLen := len(in)
	oldBytesLen := len(d.LastRemainBytes)
	if oldBytesLen != d.ReadLen {
		toIndex := math.Min(inLen, *readIndex+d.ReadLen-oldBytesLen)
		d.LastRemainBytes = append(d.LastRemainBytes, in[*readIndex:toIndex:toIndex]...)
		*readIndex = toIndex
	}
}

func Encode(in []byte) (out []byte) {
	len := len(in)
	out = make([]byte, UseLengthHeadLen+len)
	copy(out, math.IntToBytes(len, UseLengthHeadLen))
	copy(out[UseLengthHeadLen:], in)
	return out
}
