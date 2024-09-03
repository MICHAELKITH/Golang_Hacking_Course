//Hacking websites

package main

import (
	"fmt"
	"net"
)

func main() {
	//Tcp
	//scann more ports

	//for loop

	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)

		if err == nil {
			break
		}
		conn.Close()

		fmt.Printf("%d open\n", i)
	}
}
