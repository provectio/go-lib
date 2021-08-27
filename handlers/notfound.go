package handlers

import (
	"bytes"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}

	http.Error(w, "Page Not Found </br> by Provectio", http.StatusNotFound)

	buf.WriteTo(w)
}
