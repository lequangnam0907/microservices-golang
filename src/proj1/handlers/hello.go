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
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("helloWorld")
	d, err := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Hello %s", d)
	if err != nil {
		http.Error(w, "Opps", http.StatusBadRequest)
		return
	}

}
