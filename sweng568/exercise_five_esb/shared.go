package exercise_four_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Agreed upon message structure
type Student struct {
	StudID           string
	Name             string
	SSN              string
	EmailAddress     string
	HomePhone        string
	HomeAddr         string
	LocalAddr        string
	EmergencyContact string
	ProgramID        string
	PaymentID        string
	AcademicStatus   string
}

type EsbMessage struct {
	Student Student
}

//
type SubRequest struct {
	Address string
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

// get the response body
func ResponseHandler(r *http.Response) (body []byte) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	ErrorHandler(err)
	return body
}

// get the response body
func RequestHandler(r *http.Request) (body []byte) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	ErrorHandler(err)
	return body
}
