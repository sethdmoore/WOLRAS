package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os/exec"
)

func handleConnection(c *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, err := c.Read(buffer[0:])

	if err != nil {
		fmt.Printf("Error reading packet! %s", err)
		return
	}

	if n != 102 {
		fmt.Printf("Malformed magic packet, len %d, want 102", n)
		return
	}

	//fmt.Println(hex.EncodeToString(buffer[6:n]))
	//fmt.Println("---")

	validation := hex.EncodeToString(buffer[0:6])
	fmt.Println(validation)

	macs := hex.EncodeToString(buffer[6:n])
	mac := macs[0:12]

	for i := 0; i < 192; i += 12 {
		fmt.Printf(macs[i:i+12] + "\n")
	}

	if mac == "d05099a2feaf" {
		args := []string{"start", "win10"}
		if out, err := exec.Command("virsh", args...).Output(); err != nil {
			fmt.Printf("Error starting VM: %s %s\n", out, err)
			return
		}
	}

	//fmt.Printf(hex.EncodeToString(buffer[0:6]))
}

func main() {
	port := ":9"
	protocol := "udp"

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		fmt.Printf("Incorrect address: %s", err)
		return
	}

	ln, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		fmt.Printf("Error listening on port 9!: %s", err)
		return
	}

	//defer ln.Close()

	for {
		handleConnection(ln)
	}

}
