package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"time"
)

// CreateSetReboot set reboot device
func CreateSetReboot() []byte {
	cmd := make([]byte, 4)
	//cmd[0] = 0xa0
	cmd[0] = 0x40
	cmd[1] = 0x02
	cmd[2] = 0x0E
	cmd[3] = 0xB0
	return cmd
}

// CreateSetAntena set Antena
func CreateSetAntena(ant byte) []byte {
	cmd := make([]byte, 5)
	cmd[0] = 0x40
	cmd[1] = 0x03
	cmd[2] = 0x0A
	cmd[3] = ant
	cmd[4] = CalcCheckSum(cmd)
	return cmd
}

// CreateReadTagID Read tag id
func CreateReadTagID() []byte {
	cmd := make([]byte, 8)
	cmd[0] = 0x40
	cmd[1] = 6
	cmd[2] = 0xEE
	cmd[3] = 1
	cmd[4] = 0x0
	cmd[5] = 0x0
	cmd[6] = 0x0
	cmd[7] = CalcCheckSum(cmd)
	return cmd
}

// CalcCheckSum calc checksum all byte 64 3 10 0 179 = 40 3 A 0 B3
func CalcCheckSum(buf []byte) byte {
	var csum byte
	for _, b := range buf {
		csum = csum + b
	}
	csum = (^csum) + 1
	return csum
}

// CheckError error checked
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// RequestReadMsg request Message
func RequestReadMsg(conn *net.UDPConn) {
	cmd := CreateReadTagID()
	_, err := conn.Write(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

// RequestSetAnt request set Antena
func RequestSetAnt(conn *net.UDPConn) {
	cmd := CreateReadTagID()
	_, err := conn.Write(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

// listen listen read from replay
func listen(connection *net.UDPConn, quit chan struct{}) {
	buffer := make([]byte, 30)
	var req MiTagReq

	_, _, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		_, _, err = connection.ReadFromUDP(buffer)
		if err == nil {
			if buffer[0] == 0xF0 {
				s := ""
				n := int(buffer[1])
				for i := 5; i <= n; i++ {
					s += fmt.Sprintf("%02x", buffer[i])
				}
				req.Encode = s
				req.MiTag()
			}
			time.Sleep(time.Millisecond * 125)
			go RequestReadMsg(connection)
		}
	}
	quit <- struct{}{}
}

func main() {

	ServerAddr, err := net.ResolveUDPAddr("udp", "192.168.1.250:1969")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", ":5000")
	CheckError(err)

	conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	defer conn.Close()

	RequestSetAnt(conn)
	quit := make(chan struct{})
	for i := 0; i < runtime.NumCPU(); i++ {
		listen(conn, quit)
	}
	<-quit
}
