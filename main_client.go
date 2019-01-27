package main

import (
	"flag"
	"fmt"
)

func main_client() {
	// Analyse the command line.
	host := flag.String("h", "localhost", "TCP server host")
	port := flag.Int("p", Port, "TCP server port")
	flag.Parse()
	Port = *port
	Host := *host
	fmt.Println("The time server address: ", GetAddress(Host))

	// Start the local client.
	client := Client{}
	fmt.Println(client.Request(Host))
}
