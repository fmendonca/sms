package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/fmendonca/sms/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todasEntradas := models.BuscaItem()
	temp.ExecuteTemplate(w, "Index", todasEntradas)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		telefone := r.FormValue("telefone")
		status := r.FormValue("status-select")
		models.CriarItem(nome, telefone, status)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idCad := r.URL.Query().Get("id")
	models.DeletaItem(idCad)
	log.Println("Id solicitado para deletar", idCad)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idCad := r.URL.Query().Get("id")
	cadastro := models.EditItem(idCad)
	temp.ExecuteTemplate(w, "Edit", cadastro)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		telefone := r.FormValue("telefone")
		status := r.FormValue("status-select")

		idConvert, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro de conversao do ID para o int:", err)
		}
		models.AtualizarItem(idConvert, nome, telefone, status)
	}
	http.Redirect(w, r, "/", 301)

}

func CriaMSGSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		smsmensagem := r.FormValue("mensagem-sms")
		models.CadastrarMSG(smsmensagem)
	}
	http.Redirect(w, r, "/", 301)
}
