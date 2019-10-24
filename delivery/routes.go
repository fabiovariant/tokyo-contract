package delivery

import (
	"net/http"

	"github.com/fabiovariant/tokyo-domains/route"
	"github.com/gorilla/mux"
)

// GetRoutes get API routes
func getRoutes(d ClientContractsDelivery) (r route.Routes) {
	r = route.Routes{
		route.Route{
			Name:        "NewContract",
			Method:      "POST",
			Pattern:     "/contract",
			HandlerFunc: d.NewContract,
		},
		route.Route{
			Name:        "GetContractByID",
			Method:      "GET",
			Pattern:     "/contract/{id}",
			HandlerFunc: d.GetContractByClientID,
		},
	}
	return
}

// GetMuxRoutes returns all internal routes for house service
func GetMuxRoutes(cd ClientContractsDelivery) (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(false)
	for _, route := range getRoutes(cd) {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return
}
