package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Create simple tcp socket
	listener, err := net.Listen("tcp", "localhost:8080") // Listen tcp protocol in localhost at port 8080
	if err != nil {
		fmt.Println(err) // Print error if there is some error
		os.Exit(1)       // immediately close the app if there is error in listener
	}

	// forever loop. Will looping forever unless the app is closed or stopped
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue // if there is error then we will continue to the next loop
		}

		// Write out the accepted connection time with anonymous inline function
		go func() { // we can add `go` to make this function as goroutine so this will not blocked accept connection
			for {
				_, err := io.WriteString(conn, time.Now().Format("15:04:05\n")) // Write connection time
				if err != nil {
					return // probably there is error so we return nothing
				}
				time.Sleep(time.Second) // this will printout time every 1 second
			}
		}()
	}
}
