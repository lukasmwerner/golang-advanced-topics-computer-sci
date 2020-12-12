package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Status the basic status code
type Status struct {
	StatusCode    int    `json:"status"`
	StatusMessage string `json:"reason,omitempty"`
}

// WriteResponse writes the status response as json to the io writer
func (s *Status) WriteResponse(w http.ResponseWriter) (*Status, error) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(s)
	if err != nil {
		return s, fmt.Errorf("Status-WriteResponse: ", err)
	}
	_, err = w.Write(b)
	if err != nil {
		return s, fmt.Errorf("Status-WriteResponse: ", err)
	}
	return s, nil
}

// SuccessStatus is a status that is sucessful
func SuccessStatus() *Status {
	return &Status{StatusCode: 200, StatusMessage: "OK"}
}

// StateSuccessStatus returns a status with the given message
func StateSuccessStatus(status string) *Status {
	return &Status{StatusCode: 200, StatusMessage: status}
}

// ErrorStatus Creates a status error from a given error
func ErrorStatus(err error) *Status {
	return &Status{StatusCode: 500, StatusMessage: err.Error()}
}

// ContainerRequest type
type ContainerRequest struct {
	Image string `json:"image"`
}

type Response struct {
	Status *Status     `json:"status"`
	Body   interface{} `json:"data"`
}

func (r Response) WriteResponse(w http.ResponseWriter) (Response, error) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(r)
	if err != nil {
		return r, fmt.Errorf("Status-WriteResponse: ", err)
	}
	_, err = w.Write(b)
	if err != nil {
		return r, fmt.Errorf("Status-WriteResponse: ", err)
	}
	return r, nil
}
