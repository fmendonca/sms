package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fmendonca/sms/routes"
)

func main() {

	port := map[bool]string{true: os.Getenv("PORT"), false: "8000"}[os.Getenv("PORT") != ""]

	routes.Rotas()
	log.Println("Server is listening at", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}
