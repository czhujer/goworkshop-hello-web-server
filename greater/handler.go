package greater

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"log"
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

type addRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type addResponse struct {
	C int `json:"c"`
}

func (g *Handler) Add(rw http.ResponseWriter, req *http.Request) {
	var requestBody addRequest
	var response addResponse

	err := json.NewDecoder(req.Body).Decode(&requestBody)

	if err != nil {
		http.Error(rw, "unable to parse json", http.StatusBadRequest)
		log.Print("bad request: ", err)
		return
	}

	spew.Dump(&requestBody)

	response.C = requestBody.A + requestBody.B

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		panic(err)
	}
}
