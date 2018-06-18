package main

import (
	"log"
	"net/http"

	"github.com/sgeffken/cdays-go/internal/routing"
)

func main() {
	log.Printf("The application is starting and listening to port 8000")
	/*
		version.Release = "1"
		version.BuildTime = "now"
		version.Commit = "00000"
		log.Printf("The application is starting, version is %s, build time is %s, commit is %s...", version.Release, version.BuildTime, version.Commit)
	*/
	/*
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
			//		log.Fatal("No port provided")
		}
	*/
	r := routing.NewDiagnosticsRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
