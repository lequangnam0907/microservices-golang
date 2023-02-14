package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}
func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Bye")
	d, error := ioutil.ReadAll(r.Body)
	if error != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Goodbye %s", d)
}
