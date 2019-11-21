package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"goworkshop-hello-web-server/greater"
	"net/http"
)

func main() {
	dottedstring := `.
..
...
`
	greaterHandler := &greater.Handler{}

	router := chi.NewRouter()
	router.Get("/hello", greaterHandler.Hello)
	router.Get("/bye", greaterHandler.Bye)

	spew.Dump(greaterHandler, "blauh")

	fmt.Println("Starting stupid web server..")
	fmt.Println(dottedstring)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
