package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	// Analyse the command line.
	port := flag.Int("p", Port, "TCP server port")
	flag.Parse()
	Port = *port
	fmt.Println("The time server address: ", GetAddress("localhost"))

	// Maximize CPU usage for the maximal performance.
	cores := runtime.NumCPU()
	//fmt.Printf("This machine has %d CPU cores. \n", cores)
	runtime.GOMAXPROCS(cores)

	// The server state initialization.
	server := Server{
		isRunning: true,
		listener:  nil,
	}
	// Start the server.
	server.Run()

	fmt.Println("Server was terminated.")
}
