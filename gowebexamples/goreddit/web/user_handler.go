package web

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
	"mriyam.com/goreddit"
)

type UserHandler struct {
	store    goreddit.Store
	sessions *scs.SessionManager
}

func (u *UserHandler) Register() http.HandlerFunc {
	type data struct {
		SessionData
		CSRF template.HTML
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/user_register.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data{
			SessionData: GetSessionData(u.sessions, r.Context()),
			CSRF:        csrf.TemplateField(r),
		})
	}
}

func (u *UserHandler) RegisterSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := RegisterForm{
			Username:      r.FormValue("username"),
			Password:      r.FormValue("password"),
			UsernameTaken: false,
		}

		if _, err := u.store.UserByUsername(form.Username); err == nil {
			form.UsernameTaken = true
		}

		if !form.Validate() {
			u.sessions.Put(r.Context(), "form", form)
			http.Redirect(w, r, r.Referer(), http.StatusFound)
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := u.store.CreateUser(&goreddit.User{
			ID:       uuid.New(),
			Username: form.Username,
			Password: string(password),
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u.sessions.Put(r.Context(), "flash", "Your registration was successful. Please login.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (u *UserHandler) Login() http.HandlerFunc {
	type data struct {
		SessionData
		CSRF template.HTML
	}
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/user_login.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data{
			SessionData: GetSessionData(u.sessions, r.Context()),
			CSRF:        csrf.TemplateField(r),
		})
	}
}

func (u *UserHandler) LoginSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := LoginForm{
			Username:             r.FormValue("username"),
			Password:             r.FormValue("password"),
			IncorrectCredentials: false,
		}

		user, err := u.store.UserByUsername(form.Username)
		if err != nil {
			form.IncorrectCredentials = true
		} else {
			compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
			form.IncorrectCredentials = compareErr != nil
		}

		if !form.Validate() {
			u.sessions.Put(r.Context(), "form", form)
			http.Redirect(w, r, r.Referer(), http.StatusFound)
			return
		}

		u.sessions.Put(r.Context(), "user_id", user.ID)
		u.sessions.Put(r.Context(), "flash", "You have been logged in successfully.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (u *UserHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u.sessions.Remove(r.Context(), "user_id")
		u.sessions.Put(r.Context(), "flash", "You have been logged out successfully.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
