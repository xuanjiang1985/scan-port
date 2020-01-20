package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	startPort = 20
	endPort   = 100
)

var (
	openedPort = make([]int, 0, 100)
	timeout    = time.Duration(time.Millisecond * 300)
)

func main() {
	for i := startPort; i < endPort; i++ {
		_, err := net.DialTimeout("tcp", "www.shangke.ltd:"+strconv.Itoa(i), timeout)
		if err != nil {
			fmt.Println(err)
			continue
		}
		openedPort = append(openedPort, i)
	}
	fmt.Println("opened ports: ", openedPort)
}
