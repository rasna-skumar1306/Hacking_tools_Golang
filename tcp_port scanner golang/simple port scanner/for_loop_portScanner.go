package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 30; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open \n", i)
	}
}
