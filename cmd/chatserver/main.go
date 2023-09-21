package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./img/Go-Logo_Fuchsia.svg")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/index.html")
}

func main() {
	router := mux.NewRouter()
	middlewares := negroni.Classic()
	router.HandleFunc("/favicon.ico", faviconHandler)
	router.HandleFunc("/", homeHandler).Methods("GET")

	middlewares.UseHandler(router)
	http.ListenAndServe(":48391", middlewares)
}
