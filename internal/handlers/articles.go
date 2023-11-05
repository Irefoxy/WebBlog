package handlers

import (
	"WebBlog/internal/model"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) articleHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора статьи из URL
	vars := mux.Vars(r)
	articleID := vars["id"]

	// Получение содержимого статьи из базы данных
	article, err := h.getArticle(articleID)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Отображение статьи на странице
	renderTemplate(w, "internal/templates/article.html", article)
}

func (h *Handler) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	if !h.checkLoginState(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	err := h.createArticle(title, content)
	if err != nil {
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) getArticles(offset int) ([]model.Article, int, error) {
	var dbArticles []model.DbArticle
	var total int64
	var articles []model.Article

	if err := h.DB.Order("id asc").Offset(offset).Limit(3).Find(&dbArticles).Error; err != nil {
		return nil, 0, err
	}

	// Вычисление общего количества статей в базе данных для пагинации
	h.DB.Model(&model.DbArticle{}).Count(&total)

	// Преобразование содержимого статей из Markdown в HTML
	for _, dbArt := range dbArticles {
		article := model.Article{
			ID:      dbArt.ID,
			Title:   dbArt.Title,
			Content: template.HTML(blackfriday.Run([]byte(dbArt.Content))),
		}
		articles = append(articles, article)
	}

	return articles, int(total), nil
}

// getArticle получает одну статью по ID
func (h *Handler) getArticle(id string) (model.Article, error) {
	var dbArt model.DbArticle // Используйте тип dbArticle из вашего пакета models

	// Поиск статьи по id
	if err := h.DB.Where("id = ?", id).First(&dbArt).Error; err != nil {
		return model.Article{}, err
	}

	// Преобразование содержимого статьи из Markdown в HTML
	article := model.Article{
		ID:      dbArt.ID, // Используйте Id для dbArticle и ID для Article
		Title:   dbArt.Title,
		Content: template.HTML(blackfriday.Run([]byte(dbArt.Content))),
	}

	return article, nil
}

func (h *Handler) createArticle(title, content string) error {
	dbArt := model.DbArticle{ // Используйте тип dbArticle для создания записи
		Title:   title,
		Content: content,
	}

	if err := h.DB.Create(&dbArt).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
