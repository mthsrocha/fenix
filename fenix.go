package main

import (
	"fmt"

	"github.com/Fenix/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting application web...")

	server.HttpServer()
}
