package route

import (
	"casetask/hehe/src/swg"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = swg.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Привіт, світе!\nЦей API працює ;)")

}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	Route{
		"Rate",
		strings.ToUpper("Get"),
		"/api/rate",
		swg.Rate,
	},

	Route{
		"SendEmails",
		strings.ToUpper("Post"),
		"/api/sendEmails",
		swg.SendEmails,
	},

	Route{
		"Subscribe",
		strings.ToUpper("Post"),
		"/api/subscribe",
		swg.Subscribe,
	},
}
