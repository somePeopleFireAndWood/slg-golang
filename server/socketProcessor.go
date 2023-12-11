package server

import (
	"bufio"
	"fmt"
	"net"
	"slg-golang/server/decoder"
)

func ProcessConn(conn net.Conn) {
	go func() {
		channel := Channel{
			InBuf:         [1024]byte{},
			Conn:          &conn,
			LengthDecoder: decoder.NewLengthDecoder(),
		}
		channel.process()
	}()
}

type Channel struct {
	InBuf         [1024]byte
	Conn          *net.Conn
	LengthDecoder *decoder.LengthDecoder
}

func (c *Channel) process() {
	conn := *(c.Conn)
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		bufArr := c.InBuf
		n, err := reader.Read(bufArr[:]) // 读取数据
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			break
		}

		outArr := c.LengthDecoder.Decode(bufArr[:n])
		for _, singleOut := range outArr {
			recvStr := string(singleOut)
			fmt.Println("收到client端发来的数据：", recvStr)
			_, err := conn.Write([]byte(recvStr))
			if err != nil {
				fmt.Println("read from conn failed, err:", err)
			} // 发送数据
		}
	}
}
