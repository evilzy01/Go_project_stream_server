package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
