package server

import (
	"net/http"
	"github.com/Fenix/internal/handlers"
)

func HttpServer() {
	handlers.HomeHandlers()
	handlers.Handlers()
	http.ListenAndServe(":8000", nil)
}
