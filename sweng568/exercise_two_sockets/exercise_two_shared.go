package exercise_two_sockets

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"
)

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

// Prints a string of a message and the time
func LogMessage(message string) {
	t := time.Now().Local().Format("2018-10-24 00:00:00")
	fmt.Println(t + ": " + message)
}

func ErrorHandler(e error) {
	if e != nil {
		LogMessage("error - " + e.Error())
		os.Exit(1)
	}
}

func DigestMessage(b []byte) (hash []byte) {
	hasher := sha256.New()
	hasher.Write(b)
	hash = hasher.Sum(nil)
	return
}
