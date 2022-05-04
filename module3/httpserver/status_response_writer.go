package main

import (
	"net/http"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *StatusResponseWriter) Status(statusCode int) {
	w.WriteHeader(statusCode)
	w.statusCode = statusCode
}
