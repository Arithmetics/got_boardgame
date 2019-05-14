package app

import (
	"net/http"

	u "github.com/arithmetics/got_boardgame/utils"
)

// NotFoundHandler responds to bad requests
var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}
