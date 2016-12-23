package thebrief

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func articleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	content := map[string]string{"Category": ps.ByName("category"), "Slug": ps.ByName("slug")}
	err := renderWithPartials(w, "article", "Article Page", content)

	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
