/* Create a basic server using TCP

- The server should use net.Listen ot listen on port 8080.
- Remember to close the listener using `defer`
- Remember that from the "net" package, you first need to Listen,
then you need to Accept an incoming connection
- Write a response back on the connection
- Use io.WriteString to write the response: I see you connected.
- Remember to close the connection
- Once you have all of that working, run your TCP server and test it
from telnet (telnet localhost 8080)

*/
package main

import (
	"io"
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
		io.WriteString(conn, "I see you connected.\n")
		conn.Close()
	}
}
