package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting cloudpc on :8080")
	http.HandleFunc("/api/v1/container", ContainerCreationHandler)
	http.HandleFunc("/api/v1/containers/list", ContainerListHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
