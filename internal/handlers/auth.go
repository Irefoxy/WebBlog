package handlers

import (
	"log"
	"net/http"
)

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "internal/templates/creds.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == h.Creds.User && password == h.Creds.Password {
		session, err := h.SessionStore.Get(r, "session-auth")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["loggedIn"] = true
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusFound)
	} else {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}
}

func (h *Handler) checkLoginState(r *http.Request) bool {
	session, err := h.SessionStore.Get(r, "session-auth")
	if err != nil {
		log.Println(err)
		return false
	}
	value := session.Values["loggedIn"]
	if value == nil {
		return false
	}

	return value.(bool)
}
