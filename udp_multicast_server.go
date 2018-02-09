package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Start tick server at 224.0.0.1:9999")

	//  IPv4 先頭4ビットが1110のアドレス（224.0.0.0 ～ 239.255.255.255）がマルチキャストアドレス
	conn, err := net.Dial("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start := time.Now()
	wait := 10*time.Second - time.Nanosecond*time.Duration(start.UnixNano()%(10*1000*1000*1000))
	time.Sleep(wait)

	ticker := time.Tick(10 * time.Second)
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
