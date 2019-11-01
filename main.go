package main

import (
	"fmt"
	"net/http"
)

type Greater struct {
}

type Router struct {
	greater *Greater
}

func (g *Greater) Hello(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusTeapot)
	rw.Write([]byte("I'm teapot\n"))
}

func (g *Greater) Bye(rw http.ResponseWriter, req *http.Request) {
	//rw.WriteHeader(http.StatusTeapot)
	rw.Write([]byte("Bye\n"))
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/hello":
		r.greater.Hello(rw,req)
	case "/bye":
		r.greater.Bye(rw, req)
	default:
		http.NotFound(rw, req)
	}
}

func main() {
	dottedstring := `.
..
...
`
	greater := &Greater{}

	router := &Router{
		greater: greater,
	}

	fmt.Println("Starting web server..")
	fmt.Println(dottedstring)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
