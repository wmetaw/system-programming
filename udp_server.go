package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8888")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		// ReadFrom()で通信内容を読み込み、クライアントのアドレス情報を受け取る
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v:%v\n", remoteAddress, string(buffer[:length]))

		// ReadFrom()で取得したアドレスに対しては、net.PacketConnインタフェースのWriteTo()メソッドを使ってデータを返送することができる
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
