package thebrief

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router := httprouter.New()

	router.GET("/:category/:slug", articleHandler)
	router.GET("/", homeHandler)

	http.Handle("/", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := renderWithPartials(w, "home", "Welcome", nil)

	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
