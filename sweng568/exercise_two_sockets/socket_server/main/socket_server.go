package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

// James Giltner
// SWENG568
// Exercise 2: Sockets
// 10/29/2018

// TCP sockets server written to share a student record between the Roster
// System and the Course System

type Student struct {
	StudID           int
	Name             string
	SSN              string
	EmailAddress     string
	HomePhone        string
	HomeAddr         string
	LocalAddr        string
	EmergencyContact string
	ProgramID        int
	PaymentID        string
	AcademicStatus   int
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Service is listening on localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		buffer := make([]byte, 4)
		conn.Read(buffer)
		request := string(buffer)
		fmt.Println(request)
		if request == "1111" {
			student := Student{
				StudID:           1111,
				Name:             "Bob Smith",
				SSN:              "222-333-1111",
				EmailAddress:     "bsmith@yahoo.com",
				HomePhone:        "215-777-8888",
				HomeAddr:         "123 Tulip Road, Ambler, PA 19002",
				LocalAddr:        "321 Maple Avenue, Lion Town, PA 16800",
				EmergencyContact: "John Smith (215-222-6666)",
				ProgramID:        206,
				PaymentID:        "1111-206",
				AcademicStatus:   1,
			}
			//ar response *[]byte
			response, _ := json.Marshal(student)
			fmt.Println(string(response) + "\n")
			conn.Write(response)
		} else {
			conn.Write([]byte("null\n"))
		}
		conn.Close()
	}
}
