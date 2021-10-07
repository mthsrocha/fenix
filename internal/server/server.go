package server

import (
	"net/http"
	"github.com/Fenix/internal/handlers"
)

type Server struct {

}


func HttpServer() {
	handlers.HomeHandlers()
	handlers.AuthHandlers()
	handlers.UserHandlers()
	handlers.Handlers()
	http.ListenAndServe(":8000", nil)
}
