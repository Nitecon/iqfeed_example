package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
	"io"

)

var (
	filename = "../testdata.dat"
)

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
}


func main() {
	service := "127.0.0.1:5009"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Printf("I'm listening on: %s\n", tcpAddr.String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}

		handleClient(conn)
	}
}

// Read function does as expected and reads data from the network stream.
func HandleReading(Conn net.Conn) {
	r := bufio.NewReader(Conn)
	for {
		select {
		default:
			line, isPrefix, err := r.ReadLine()
			for err == nil && !isPrefix {
				fmt.Println(line)
				//c.ProcessReceiver(line)
				line, isPrefix, err = r.ReadLine()
			}
			if isPrefix {
				log.Println("buffer size to small")
				//return Do not return and break the loop
			}
			if err != io.EOF {
				log.Printf("Pipe closed, exiting: %s\n...",err)
				Conn.Close()
				os.Exit(0)
			}
		}
	}

}


func handleClient(conn net.Conn) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error loading file(%s): %s", filename, err.Error())
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		d, err := r.ReadString(10) // 0x0A separator = newline
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err.Error()) // if you return error
			break
		}
		conn.Write([]byte(d))
		fmt.Printf("Data: %+v\n", d)
	}
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
