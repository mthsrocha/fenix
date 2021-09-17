package server

import (
	"net/http"
	"github.com/Fenix/internal/handlers"
)

func HttpServer() {
	handlers.Handlers()
	http.ListenAndServe(":8000", nil)
}
