package thebrief

import (
	"fmt"
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
	fmt.Fprint(w, "Hello, world!")
}
