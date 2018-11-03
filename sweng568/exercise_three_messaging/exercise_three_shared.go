// Shared tools for socket server and client:
//
package exercise_two_sockets

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"
)

// Agreed upon message structure
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

//
type Topic struct {
	Topic string
}

// Prints a string of a message and the time
//
func LogMessage(message string) {
	t := time.Now().Local().Format(time.RFC1123)
	fmt.Println(t + ": " + message)
}

// log an error if it occurs and exit
//
func ErrorHandler(e error) {
	if e != nil {
		LogMessage("error - " + e.Error())
		os.Exit(1)
	}
}

// digest a message using SHA256
//
func DigestMessage(b []byte) (hash []byte) {
	hasher := sha256.New()
	hasher.Write(b)
	hash = hasher.Sum(nil)
	return
}
