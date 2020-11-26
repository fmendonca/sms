package routes

import (
	"net/http"

	"github.com/fmendonca/sms/controllers"
)

func Rotas() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/sendsms", controllers.CriaMSGSMS)
	http.HandleFunc("/api/health", controllers.Alive)
}
