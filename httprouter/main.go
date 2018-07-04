package main

import (
	"net/http"

	bonehttprouter "github.com/ariefrahmansyah/bone-httprouter"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	for _, route := range bonehttprouter.StaticRoutes {
		router.Handle(route.Method, route.Path, httpRouterHandle)
	}

	http.ListenAndServe(":8081", router)
}

func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}
