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
		case "POST":
			mux.Post(route.Path, http.HandlerFunc(httpHandlerFunc))
		case "PUT":
			mux.Put(route.Path, http.HandlerFunc(httpHandlerFunc))
		case "PATCH":
			mux.Patch(route.Path, http.HandlerFunc(httpHandlerFunc))
		case "DELETE":
			mux.Delete(route.Path, http.HandlerFunc(httpHandlerFunc))
		default:
			panic("Unknow HTTP method: " + route.Method)
		}
	}

	http.ListenAndServe(":8080", mux)
}

func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}
