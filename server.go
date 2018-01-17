package main

import (
	"github.com/acenolaza/rest-api-sample/routers"
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := routers.InitRouter()

	log.Println("Server is now listening...")

	// handle server interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Stopping the server...")
		// add cleanup code here
		os.Exit(1)
	}()

	n := negroni.Classic()
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}
