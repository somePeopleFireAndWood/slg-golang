package server

import (
	"bufio"
	"fmt"
	"github.com/somePeopleFireAndWood/slg-golang/encoderDecoder/lengthEncoderDecoder"
	"net"
)

func ProcessConn(conn net.Conn) {
	go func() {
		channel := Channel{
			InBuf:         [1024]byte{},
			Conn:          &conn,
			LengthDecoder: lengthEncoderDecoder.NewLengthDecoder(),
		}
		channel.process()
	}()
}

type Channel struct {
	InBuf         [1024]byte
	Conn          *net.Conn
	LengthDecoder *lengthEncoderDecoder.LengthDecoder
}

func (c *Channel) process() {
	conn := *(c.Conn)
	defer conn.Close() // 关闭连接
	reader := bufio.NewReader(conn)
	for {
		bufArr := c.InBuf
		n, err := reader.Read(bufArr[:]) // 读取数据
		if err != nil {
			fmt.Println("read from conn failed, writeErr:", err)
			break
		}

		//recvStr := string(bufArr[:n])
		//fmt.Println("收到client端发来的数据：", recvStr)
		//_, writeErr := conn.Write([]byte(recvStr))
		//if writeErr != nil {
		//	fmt.Println("read from conn failed, writeErr:", writeErr)
		//} // 发送数据

		outArr := c.LengthDecoder.Decode(bufArr[:n])
		for _, singleOut := range outArr {
			recvStr := string(singleOut)
			fmt.Println("收到client端发来的数据：", recvStr)
			_, writeErr := conn.Write([]byte(recvStr))
			if writeErr != nil {
				fmt.Println("read from conn failed, writeErr:", writeErr)
			} // 发送数据
		}
	}
}
