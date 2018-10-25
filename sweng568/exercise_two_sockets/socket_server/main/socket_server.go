package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	shared "github.com/corygiltner/go-pennstate/sweng568/exercise_two_sockets"
	"github.com/google/go-cmp/cmp"
	"log"
	"net"
	"os"
)

// James Giltner
// SWENG568
// Exercise 2: Sockets
// 10/29/2018

// TCP sockets server :
// listens for single student json messages

func main() {
	shared.LogMessage("server - starting student socket integration")
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listener.Close()
	shared.LogMessage("service is listening on localhost:8888")
	for {
		conn, err := listener.Accept()
		shared.ErrorHandler(err)
		shared.LogMessage("request received")

		// create a buffer to hold the request message and read
		// the message into the buffer
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		request := bytes.Trim(buffer, "\x00")
		// log the message
		shared.LogMessage("student: " + string(request))

		// create a student object from a json message
		student := shared.Student{}
		err = json.Unmarshal(request, &student)
		shared.ErrorHandler(err)
		shared.LogMessage("saving student: " + student.Name)
		// check to make sure the object isn't empty if it is
		// respond with an empty message otherwise return a checksum
		// and log the sum
		if cmp.Equal(student, shared.Student{}) {
			shared.LogMessage("error - not able to save student")
			conn.Write([]byte(""))
		} else {
			shared.LogMessage("student record saved to course system")
			hash := shared.DigestMessage(request)
			conn.Write(hash)
			shared.LogMessage("response: " + hex.EncodeToString(hash))
		}
		shared.LogMessage("awaiting next student record")
		conn.Close()
	}
}
