package main

import (
	"log"
	"net"

	"xip/xip"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 53})
	if err != nil {
		log.Fatal(err.Error())
	}

	query := make([]byte, 512)

	for {
		_, addr, err := conn.ReadFromUDP(query)
		if err != nil {
			log.Println(err.Error())
			break
		}

		go func() {
			response, logMessage, err := xip.QueryResponse(query)
			if err != nil {
				log.Println(err.Error())
				return
			}
			_, err = conn.WriteToUDP(response, addr)
			log.Printf("%v.%d %s", addr.IP, addr.Port, logMessage)
		}()
	}
}
