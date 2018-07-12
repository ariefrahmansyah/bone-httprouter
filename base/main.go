package main

import (
	"net/http"

	bonehttprouter "github.com/ariefrahmansyah/bone-httprouter"
)

func main() {
	for _, route := range bonehttprouter.StaticRoutes {
		http.HandleFunc(route.Path, httpHandlerFunc)
	}

	http.ListenAndServe(":8080", nil)
}

func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Path))
}
