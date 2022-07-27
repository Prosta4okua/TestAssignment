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

func Index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(w, "Привіт, світе!\nЦей API працює ;)")
	if err != nil {
		return
	}

}

var routes = Routes{

	// проста перевірка чи працює API взагалі
	Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	// отримати поточний курс
	Route{
		"Rate",
		strings.ToUpper("Get"),
		"/api/rate",
		swg.Rate,
	},

	// надіслати поточний курс по скриньках
	Route{
		"SendEmails",
		strings.ToUpper("Post"),
		"/api/sendEmails",
		swg.SendEmails,
	},

	// додати нову скриньку для розсилання курсу
	Route{
		"Subscribe",
		strings.ToUpper("Post"),
		"/api/subscribe",
		swg.Subscribe,
	},
}
