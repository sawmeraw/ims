package handlers

import (
	"net/http"

	app "github.com/sawmeraw/ims"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home",
	}
	app.RenderTemplate(w, "home", data)
}
