package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

func main() {
	counter, err := NewPostgresCounter(os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Printf("Error initializing counter: %#v", err)
		os.Exit(1)
	}

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", renderHandler(counter))
	router.HandleFunc("/index.html", renderHandler(counter))
	router.HandleFunc("/counter", counterHandler(counter))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port: %s", port)

	http.ListenAndServe(":"+port, loggedRouter)
}

func renderHandler(counter *PostgresCounter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		render := render.New(render.Options{Layout: "layout"})
		n, err := counter.Count()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.HTML(w, http.StatusOK, "counter",
			map[string]string {
				"count": strconv.Itoa(n),
			})
	}
}

func counterHandler(counter *PostgresCounter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		if r.FormValue("type") == "inc" {
			err = counter.Inc()
		} else if r.FormValue("type") == "decr" {
			err = counter.Decr()
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		count, err := counter.Count()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(strconv.Itoa(count)))
	}
}
