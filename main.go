package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func main() {
	host := flag.String("host", "localhost", "Server host")
	port := flag.String("port", "6676", "Server port")
	filter := flag.String("filter", "", "Regexp for filtering")
	prefix := flag.String("prefix", ">>> ", "System message prefix")
	flag.Parse()

	log.SetPrefix(*prefix)

	address := *host + ":" + *port
	log.Println("Start listen on:", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic("Couldn't create server on " + address)
	}

	output := newOutput(os.Stdout, newFilter(*filter))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Couldn't accept connection:", err)
		}

		go handleConnection(conn, output)
	}
}

func handleConnection(conn net.Conn, output io.Writer) {
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Client was disconected")

				return
			}

			log.Println("Couldn't read message:", err)
		}

		output.Write(message)
	}
}
