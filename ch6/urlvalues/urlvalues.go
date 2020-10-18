package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Printf("lang: %s\n", m.Get("lang"))
	fmt.Printf("q: %s\n", m.Get("q"))
	fmt.Printf("item: %s\n", m.Get("item"))
	fmt.Printf("item[]: %s\n", m["item"])

	m = nil
	fmt.Printf("item: %s\n", m.Get("item"))
}
