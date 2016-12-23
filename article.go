package thebrief

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func articleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Hello, world!")
}
