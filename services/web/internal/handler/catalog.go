package handler

import (
	"html/template"
	"localloop/services/web/internal/repository"
	"net/http"
)

type CatalogHandler struct {
	templates   *template.Template
	catalogRepo repository.CatalogRepository
}

func NewCatalogHandler(catalogRepo repository.CatalogRepository) (*CatalogHandler, error) {
	tmpl, err := template.ParseFiles(
		"internal/templates/categories/list.html",
		"internal/templates/categories/form.html",
	)
	if err != nil {
		return nil, err
	}

	return &CatalogHandler{
		templates:   tmpl,
		catalogRepo: catalogRepo,
	}, nil
}

func (h *CatalogHandler) ShowCategoriesPage(w http.ResponseWriter, r *http.Request) {
	categories, err := h.catalogRepo.ListCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Categories": categories,
	}

	h.templates.ExecuteTemplate(w, "list.html", data)
}

func (h *CatalogHandler) ShowCreateCategoryPage(w http.ResponseWriter, r *http.Request) {
	parentOptions, err := h.catalogRepo.ListCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"ParentOptions": parentOptions,
	}

	h.templates.ExecuteTemplate(w, "form.html", data)
}
