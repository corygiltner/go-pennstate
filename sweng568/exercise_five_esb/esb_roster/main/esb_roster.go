// James Giltner
// SWENG568
// Exercise 5
// 11/19/18

// Roster system restful api to send xml student messages through mulesoft

package main

import (
	"bytes"
	shared "github.com/corygiltner/go-pennstate/sweng568/exercise_five_esb"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/xml"
)

func main() {
	// base url string for the mulesoft esb
	brokerBase := "http://localhost:8081/esb/"
	port := "8888"

	r := mux.NewRouter()
	// an api route is exposed for the integration to publish student records
	r.HandleFunc("/api/integration/{action}/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		topic := vars["topic"]
		action := vars["action"]
		switch r.Method {
		case http.MethodGet:
			// find and publish new records for a topic
			switch action {
			case "pub":
				switch topic {
				case "student":
					// a new student record is found in the database for bob smith
					// the client would read the record into a new student object
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
					url := brokerBase + action + "/" + topic
					body, err := xml.Marshal(student)
					shared.ErrorHandler(err)

					// the student record is published to the mulesoft in an xml message
					resp, err :=http.Post(url, "application/xml", bytes.NewBuffer(body))
					shared.ErrorHandler(err)
					msg := shared.ResponseHandler(resp)
					w.Write(msg)
				default:
					w.Write([]byte("couldn't publish " + topic))
				}
			}
		default:
			w.Write([]byte("no action found"))
		}
	})

	http.ListenAndServe(":" + port, r)
}