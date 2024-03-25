package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "Hello %s", name)
	if err != nil {
		return
	}
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Kyle")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetingHandler)))
}
