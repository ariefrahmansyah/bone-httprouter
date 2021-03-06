package main

import (
	"net/http"

	bonehttprouter "github.com/ariefrahmansyah/bone-httprouter"
	"github.com/go-zoo/bone"
)

func main() {
	mux := bone.New()

	for _, route := range bonehttprouter.StaticRoutes {
		switch route.Method {
		case "GET":
			mux.Get(route.Path, http.HandlerFunc(httpHandlerFunc))
		default:
			panic("Unknow HTTP method: " + route.Method)
		}
	}

	http.ListenAndServe(":8081", mux)
}

func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Path))
}
