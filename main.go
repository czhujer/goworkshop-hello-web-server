package main

import (
	"fmt"
	"net/http"
)

type Greater struct {
}

func (g *Greater) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusTeapot)
	rw.Write([]byte("I'm teapot\n"))

}

func main() {
	dottedstring := `.
..
...
`
	greater := &Greater{}

	fmt.Println("Starting web server..")
	fmt.Println(dottedstring)

	http.ListenAndServe(":3000", greater)
}
