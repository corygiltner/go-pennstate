// James Giltner
// SWENG568
// Exercise 5
// 11/19/18

// Course system restful api

package main

import (
	"encoding/json"
	shared "github.com/corygiltner/go-pennstate/sweng568/exercise_five_esb"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {
	// base url string for the broker
	port := "9999"
	// a map to store all students
	students := make(map[string]shared.Student)

	r := mux.NewRouter()

	// an api route is exposed for the integration to either publish or subscribe to a topic
	r.HandleFunc("/api/integration/{action}/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		topic := vars["topic"]
		action := vars["action"]

		switch r.Method {
		// a post method
		case http.MethodPost:
			switch action {
			case "pub":
				// publish to the course will save a record
				switch topic {
				case "student":
					msg := shared.RequestHandler(r)
					var message *shared.EsbMessage
					var student shared.Student
					err := json.Unmarshal(msg, &message)
					shared.ErrorHandler(err)
					students[student.StudID] = message.Student
					m := "student record saved to course system"
					w.Write([]byte(m))
				}
			}
		}
	})

	// to search for the records a client can request the record type and search by id
	r.HandleFunc("/api/search/{record}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		record := vars["record"]

		switch r.Method {
		case http.MethodGet:
			// get method will return a count of the students
			switch record {
			case "student":
				c := strconv.Itoa(len(students))
				w.Write([]byte("there are " + c + " student(s) records in the course system"))
			}
		}
	})

	http.ListenAndServe(":" + port, r)
}