package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
)

// ContainerCreationHandler handles the creation of a container
func ContainerCreationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var cont ContainerRequest
		err := json.NewDecoder(r.Body).Decode(&cont)
		if err != nil {
			ErrorStatus(fmt.Errorf("JSON decode failed: %v", err)).WriteResponse(w)
			return
		}

		id, err := CreateNewContainer(cont.Image)
		if err != nil {
			ErrorStatus(fmt.Errorf("creating container failed: %v", err)).WriteResponse(w)
			return
		}
		resp := &Response{
			Status: SuccessStatus(),
			Body: struct {
				Container string `json:"container"`
			}{
				Container: id,
			},
		}
		resp.WriteResponse(w)
		return
	}
	ErrorStatus(fmt.Errorf("wrong http method")).WriteResponse(w)
	return
}

func ContainerListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		containers, err := GetContainers(true, true, true)
		if err != nil {
			ErrorStatus(fmt.Errorf("failed to get containers: %v", err)).WriteResponse(w)
			return
		}
		resp := &Response{
			Status: SuccessStatus(),
			Body: struct {
				Containers []types.Container `json:"containers"`
			}{
				Containers: containers,
			},
		}
		resp.WriteResponse(w)
		return
	}
	ErrorStatus(fmt.Errorf("wrong http method")).WriteResponse(w)
	return
}
