package handler

import (
	"html/template"
	"localloop/services/web/internal/repository"
	"net/http"
)

type AuthHandler struct {
	templates *template.Template
	userRepo  repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) (*AuthHandler, error) {
	tmpl, err := template.ParseFiles(
		"internal/templates/login.html",
		"internal/templates/register.html",
	)
	if err != nil {
		return nil, err
	}

	return &AuthHandler{
		templates: tmpl,
		userRepo:  userRepo,
	}, nil
}

func (h *AuthHandler) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "login.html", nil)
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	token, err := h.userRepo.Login(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		w.Write([]byte(`<div class="text-red-500">` + err.Error() + `</div>`))
		return
	}

	// Set JWT token as cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.Write([]byte(`<div class="text-green-500">Login successful! Redirecting...</div>
		<script>setTimeout(() => window.location.href = "/dashboard", 1000)</script>`))
}

func (h *AuthHandler) ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "register.html", nil)
}

func (h *AuthHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	err := h.userRepo.Register(
		r.FormValue("email"),
		r.FormValue("password"),
		r.FormValue("name"),
	)
	if err != nil {
		w.Write([]byte(`<div class="text-red-500">` + err.Error() + `</div>`))
		return
	}

	w.Write([]byte(`<div class="text-green-500">Registration successful! Redirecting to login...</div>
		<script>setTimeout(() => window.location.href = "/login", 1000)</script>`))
}
