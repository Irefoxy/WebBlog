package handlers

import (
	"WebBlog/internal/model"
	"net/http"
	"strconv"
)

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	hasPrevPage := page > 0
	empty := false

	articles, total, err := h.getArticles(page * 3)
	if err != nil {
		http.Error(w, "Failed to get articles", http.StatusInternalServerError)
		return
	}
	if page == 0 && len(articles) < 1 {
		empty = true
	} else if page < 0 || len(articles) < 1 {
		http.Error(w, "Page is out of range", http.StatusInternalServerError)
		return
	}

	hasNextPage := (total-1)/(page+1)/3 > 0

	data := model.HomePageData{
		Articles:    articles,
		HasPrevPage: hasPrevPage,
		PrevPage:    page - 1,
		HasNextPage: hasNextPage,
		NextPage:    page + 1,
		Empty:       empty,
	}

	renderTemplate(w, "internal/templates/home.html", data)
}
