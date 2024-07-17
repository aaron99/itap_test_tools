package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

var ngrokAuthToken = "your ngrok authtoken"

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtoken(ngrokAuthToken),
	)
	if err != nil {
		return err
	}

	log.Println("*******************************************************")
	log.Println("INSTO-Test :", listener.URL())
	log.Println("*******************************************************")

	return http.Serve(listener, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	log.Println("")
	log.Println("--------------")
	log.Println(string(bodyBytes))
	log.Println("")

	w.WriteHeader(200)
	fmt.Fprintf(w, "<html><head></head><body>")
	fmt.Fprintf(w, "<h1><br>Hello!! INSTO!!<br>")
	fmt.Fprintf(w, "</body></html>")
}
