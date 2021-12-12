package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			if ln == "" {
				break
			}
			fmt.Println(ln)
			fmt.Fprintf(conn, "I heard: %s\n", ln)
		}
		fmt.Fprintln(conn, "Thanks for being with us!")
		conn.Close()
		fmt.Println("Code got here.")
	}
}
