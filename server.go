package goomem

import (
	"fmt"
	"net"
)

func serverStart(port int) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	tcpListen, _ := net.ListenTCP("tcp", tcpAddr)
	go func() {
		for true {
			conn, err := tcpListen.Accept()
			if err != nil {
				break
			}
			go serverHandler(conn)
		}
	}()
}

func serverHandler(conn net.Conn) {

}
