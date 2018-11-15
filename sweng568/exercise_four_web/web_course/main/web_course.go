// James Giltner
// SWENG568
// Exercise 4
// 11/12/18

// Course system restful api

package main

import (
	"bytes"
	"encoding/json"
	shared "github.com/corygiltner/go-pennstate/sweng568/exercise_four_web"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// base url string for the broker
	brokerBase := "http://localhost:5000/broker/"
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
		case http.MethodGet:
			switch action {
			case "sub":
				// if a client requests a subscription to a topic
				// the server will contact the broker and subscribe
				// to the topic and return the message from the broker
				url := brokerBase + action + "/" + topic
				sub := shared.SubRequest{Address:"localhost:" + port}
				body, err := json.Marshal(sub)
				shared.ErrorHandler(err)
				resp, err :=http.Post(url, "application/json", bytes.NewBuffer(body))
				shared.ErrorHandler(err)
				msg := shared.ResponseHandler(resp)
				w.Write(msg)
			}
		// a post method
		case http.MethodPost:
			switch action {
			case "pub":
				// publish to the course will save a record
				switch action {
				case "student":
					msg := shared.RequestHandler(r)
					var student shared.Student
					err := json.Unmarshal([]byte(msg), student)
					shared.ErrorHandler(err)
					students[string(student.StudID)] = student
				}
			}
		}
	})

	// to search for the records a client can request the record type and search by id
	r.HandleFunc("/api/search/{record}/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		record := vars["record"]
		id := vars["id"]

		switch r.Method {
		case http.MethodGet:
			// get method will return the student record in json
			switch record {
			case "student":
				student := students[id]
				body, err := json.Marshal(student)
				shared.ErrorHandler(err)
				w.Write([]byte(body))
			}
		}

	})

	http.ListenAndServe(":9999", r)
}