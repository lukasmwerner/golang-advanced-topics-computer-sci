# An Adventure in go!

## Overview
The main project I was consistently working on this semester was my CloudPC project. This allows the user to have access to desktop apps in the browser. It uses docker containers with a `tigervncserver`  that gets dumped to a unix socket that then gets switched to a websocket that is then read in the browser with a `noVNC` client. The backend server is written in `go` for maximum concurrent users because of go channels. 

The main backend was originally written in `net/http` which is the default http server "library" built-in to go. However many things in the route handlers were quite frustrating which made me switch to `github.com/labstack/echo` one of the top libraries for golang http servers. The authentication is done with a `middleware.BasicAuth` which allows for making an authentication handler before the request is sent to the actual request handler. 

The routing is handled with JIT subdomain allocation through a global map singleton through a package I wrote which creates network proxies on each subdomain based on the container hostname.

Each route handler runs my high level container wrapper functions. Because docker is already written in golang writing the api was quite simple with it's godoc page. 
## Key Accomplishments
* Learning a new programming language
  * Learned the language and its unique features through [The Go Programming Language](http://www.gopl.io/) a wonderful book written with Alan A. A. Donovan and Brian W. Kernighan.
* Learning how to write `goroutines` that run concurrently
  * Go does coroutines or as they are called in go, `goroutines` quite effortlessly. All one needs to do is put the word `go` in front of the function call. Typically `goroutines` communicate over `channels` which one uses to return the data. Here is an example from ch8/pipeline3 
  * The code below creates two goroutines that use output channels with a `chan` data type of `int` and one locking function that reads the output of squares and prints them out. The counter is making the numbers for the squarer which is then outputting them later. This is what we call a pipeline in go
```go
package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)

	printer(squares)
}
```  
* Developing CloudPC
  * This project is described in length in the overview.
* GraphQL
  * This isn't something that I submitted any work on but it was quite an interesting api query style that I learned. The knowledge of graphql (however in python) was applied to the robotics digital-engineer project as well as a demo of graphql and flask (one of the most popular http micro-frameworks) working together.
    * [LaSalleRobots/Digital-Engineer](https://github.com/LaSalleRobots/Digital-Engineer)
    * [lwerner-lshigh/flask-graphql-demo](https://github.com/lwerner-lshigh/flask-graphql-demo)

## Summary 
I **LOVED** working with golang and echo so my future web projects will be written in golang because of its typing, coroutines, and speed. I also want to use GraphQL with many of my projects because it simplifies defining my apis with how i want to use them. I would probably have made GraphQL one of my topics because there is a lot more to cover in that technology and how to implement it. 

I will leave a link to the demo page of the cloudpc website in a submission comment.