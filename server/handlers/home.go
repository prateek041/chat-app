package handlers

import (
	"log"
	"net/http"
)

type HomePage struct {
	l *log.Logger
}

func NewHomePage() *HomePage {
	return &HomePage{}
}

func (hp *HomePage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("This is the homepage handler ", r.URL.Path)
}
