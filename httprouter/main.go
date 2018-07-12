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

	http.ListenAndServe(":8082", router)
}

func httpRouterHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Path))
}
