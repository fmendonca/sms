package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fmendonca/sms/controllers"
	"github.com/fmendonca/sms/routes"
)

func main() {

	port := map[bool]string{true: os.Getenv("PORT"), false: "8000"}[os.Getenv("PORT") != ""]
	apiAlive := http.NewServeMux()
	apiAlive.HandleFunc("/api/health", controllers.Alive)
	routes.Rotas()
	log.Println("Server is listening at", port)

	go func() {
		log.Println("Server is listening at 8001 /api/health")
		http.ListenAndServe("localhost:8001", apiAlive)
	}()

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
