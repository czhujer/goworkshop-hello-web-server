package greater

import (
	"net/http"
)

type Handler struct {
}

func (g *Handler) Hello(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusTeapot)
	rw.Write([]byte("I'm teapot\n"))
}

func (g *Handler) Bye(rw http.ResponseWriter, req *http.Request) {
	//rw.WriteHeader(http.StatusTeapot)
	rw.Write([]byte("Bye\n"))
}
