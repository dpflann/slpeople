package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	app "github.com/slpeople/app"
	chars "github.com/slpeople/characters"
	dupes "github.com/slpeople/duplicates"
	slapi "github.com/slpeople/salesloftapi"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

const (
	// TODO: Make this configurable
	salesLoftApiURL = "https://api.salesloft.com/v2/people.json"
)

var (
	apikey = flag.String("apikey", "", "SalesLoft API Key for communications with SalesLoft API (https://developers.salesloft.com/api.html)")
	port   = flag.String("port", "3000", "The port for the service. The default value is 3000.")
)

func main() {
	// Parse the flags and validate their input.
	flag.Parse()
	if *apikey == "" {
		fmt.Fprintf(os.Stderr, "An API Key is required. Please obtain a SalesLoft API key from your account or contact SalesLoft Support (support@salesloft.com) for assistance.")
		os.Exit(1)
	} else {
		log.Printf("Using API key: %s\n", *apikey)
	}
	if *port == "" {
		fmt.Fprintf(os.Stderr, "The port was set to empty string. :(")
		os.Exit(2)
	} else {
		log.Printf("Using port: %s\n", *port)
	}

	// Create the router, setup middleware, and establish routes and handlers.
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	slapi.InitializeClient(*apikey, salesLoftApiURL)
	r.Route("/people", func(r chi.Router) {
		r.Get("/", slapi.ListPeopleHandler)
		r.Get("/emails/char-frequencies", chars.EmailCharacterFrequenciesHandler)
		r.Get("/emails/duplicates", dupes.PossibleDuplicateEmailsHandler)
	})

	// Add file serving for the site's main page and other static assets.
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	app.FileServer(r, "/static", http.Dir(filesDir))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":"+*port, r)
}
