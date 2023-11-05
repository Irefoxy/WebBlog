package main

import (
	"WebBlog/internal/auth"
	dbase "WebBlog/internal/db"
	"WebBlog/internal/handlers"
	mw "WebBlog/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

func main() {
	dbs, err := dbase.Connect("./config/db.yaml")
	if err != nil {
		log.Fatalf("error in connecting to db: %v", err)
	}

	creds, err := auth.GetCreds("./config/admin.yaml")
	if err != nil {
		log.Fatalf("error in reading creds: %v", err)
	}

	handler := handlers.Handler{
		SessionStore: sessions.NewCookieStore([]byte("secret-key")),
		DB:           dbs,
		Creds:        creds,
	}

	router := mux.NewRouter()
	handler.InitializeRoutes(router)

	log.Println("Server started on port 8888")
	log.Fatal(http.ListenAndServe(":8888", mw.Limit(router)))
}
