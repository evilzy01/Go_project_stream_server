package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
