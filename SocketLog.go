package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	var host string
	var port uint
	flag.StringVar(&host, "host" ,"0.0.0.0", "host ip")
	flag.UintVar(&port, "port", 8989, "listen port")
	flag.Parse()
	log.Printf("SocketLog server start udp://%s:%d\n", host, port)
	addr := &net.UDPAddr{IP: net.ParseIP(host), Port: int(port)}
	socket, err1 := net.ListenUDP("udp4", addr)
	if err1 != nil {
		log.Println(err1)
		return
	}
	defer socket.Close()
	for {
		data := make([]byte, 4096)
		n, _, err2 := socket.ReadFromUDP(data[:])
		if err2 != nil {
			log.Println(err2)
			continue
		}
		log.Print(string(data[:n]))
	}
}
