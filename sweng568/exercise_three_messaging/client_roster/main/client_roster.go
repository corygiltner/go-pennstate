// TCP Socket client:
// sends new student records in a json message to the course system listening
// on localhost:8888
//
package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	shared "github.com/corygiltner/go-pennstate/sweng568/exercise_two_sockets"
	"github.com/google/go-cmp/cmp"
	"net"
)

func main() {
	// print messages the client is starting and act like a new record is found
	shared.LogMessage("client - starting student socket integration")
	address := "localhost:8888"

	conn, err := net.Dial("tcp", address)
	shared.ErrorHandler(err)

	shared.LogMessage("connection established at " + address)
	shared.LogMessage("found one new student record")

	// a new student record is found in the database for bob smith
	// the client would read the record into a new student object
	student := shared.Student{
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

	// the student is mashalled into a json object to be sent to the course
	// system located at "localhost:8888"
	shared.LogMessage("sending student to course system")
	request, err := json.Marshal(student)
	shared.ErrorHandler(err)

	// digest the message with sha256 hash
	hash := shared.DigestMessage(request)
	shared.LogMessage("student: " + string(request))

	// send the message
	_, err = conn.Write(request)
	shared.ErrorHandler(err)

	//
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	response := bytes.Trim(buffer, "\x00")
	shared.LogMessage("response: " + hex.EncodeToString(response))

	if cmp.Equal(response, []byte("")) {
		shared.LogMessage("error - student didn't save to the course system")
	}
	shared.LogMessage("request: " + hex.EncodeToString(hash))

	if cmp.Equal(hash, response) {
		shared.LogMessage("student record saved to course system")
	} else {
		shared.LogMessage("error - check of records failed")
	}
	shared.LogMessage("client exiting")
	conn.Close()
}
