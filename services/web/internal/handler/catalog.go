package handler

import (
	"errors"
	"html/template"
	"localloop/libs/pkg/errorbuilder"
	"localloop/services/web/internal/repository"
	apperror "localloop/services/web/internal/shared/error"
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
		var customErr *errorbuilder.CustomError
		if errors.As(err, &customErr) {
			switch customErr.Code {
			case errorbuilder.ErrNotFound:
				http.Error(w, "Categories not found", http.StatusNotFound)
			case errorbuilder.ErrInternal:
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			default:
				http.Error(w, customErr.Error(), int(customErr.Code))
			}
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Categories": categories,
	}

	if err := h.templates.ExecuteTemplate(w, "list.html", data); err != nil {
		http.Error(w, apperror.ErrTemplateRender(
			errorbuilder.WithOriginal(err),
		).Error(), http.StatusInternalServerError)
	}
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
