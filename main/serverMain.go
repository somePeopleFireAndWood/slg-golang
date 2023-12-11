package main

import (
	"fmt"
	"net"
	"slg-golang/server"
)

func ListenPort() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		server.ProcessConn(conn)
	}
}

func main() {
	ListenPort()
}
