package main

import (
	"encoding/json"
	"encoding/xml"
	shared "github.com/corygiltner/go-pennstate/sweng568/exercise_five_esb"
	"testing"
)

func TestMain_Xml(t *testing.T) {
	student := shared.Student{
		StudID:           "1111",
		Name:             "Bob Smith",
		SSN:              "222-333-1111",
		EmailAddress:     "bsmith@yahoo.com",
		HomePhone:        "215-777-8888",
		HomeAddr:         "123 Tulip Road, Ambler, PA 19002",
		LocalAddr:        "321 Maple Avenue, Lion Town, PA 16800",
		EmergencyContact: "John Smith (215-222-6666)",
		ProgramID:        "206",
		PaymentID:        "1111-206",
		AcademicStatus:   "1",
	}
	body, _ := xml.Marshal(student)
	print(string(body))

	print("\n\n\n")

	esb := shared.EsbMessage{Student:student}
	m, _ := json.Marshal(esb)
	print(string(m))

	var s shared.EsbMessage
	err := json.Unmarshal([]byte(m), s)
	print(err)

}


