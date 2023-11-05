package handlers

import (
	"WebBlog/internal/model"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

type Handler struct {
	SessionStore *sessions.CookieStore
	DB           *gorm.DB
	Creds        *model.Creds
}

func (h *Handler) InitializeRoutes(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", h.homeHandler).Methods("GET")
	router.HandleFunc("/article/{id}", h.articleHandler).Methods("GET")
	router.HandleFunc("/login", h.loginHandler)
	router.HandleFunc("/admin", h.adminHandler).Methods("GET")
	router.HandleFunc("/admin", h.createArticleHandler).Methods("POST")
}

func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
