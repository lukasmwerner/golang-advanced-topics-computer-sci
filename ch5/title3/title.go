package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panick
			err = fmt.Errorf("multiple title elemnets")
		default:
			panic(p)
		}
	}()
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("err: %v", err)
		}
		defer resp.Body.Close()
		body, err := html.Parse(resp.Body)
		if err != nil {
			log.Printf("err: %v", err)
		}
		soleTitle(body)
	}
}
