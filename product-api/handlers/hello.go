package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	 //below logger is sent back to the user
	h.l.Println("Hello World") //did just to have a control over the
	 //logger. Can be helpful if used with Dependency Injection 
	 //connetcing to database and writing unit testes

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Ooops!"))
		return
	}

	log.Printf("Hello %s\n", data) //not sent back to the user

	//sending data back to the user
	fmt.Fprintf(rw, "Hello %s", data)
}
