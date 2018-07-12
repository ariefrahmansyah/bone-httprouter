package main

import (
	"net/http"

	bonehttprouter "github.com/ariefrahmansyah/bone-httprouter"
	"github.com/ariefrahmansyah/server/router"
)

func main() {
	router := router.New()

	for _, route := range bonehttprouter.StaticRoutes {
		router.Get(route.Path, httpHandlerFunc)
	}

	http.ListenAndServe(":8083", router)
}

func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Path))
}
