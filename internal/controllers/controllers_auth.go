package controllers

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/Fenix/internal/models"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const sessionName = "sessionid"

var CookieLen = 32
var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

func CookieGenerator() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	generator := make([]byte, CookieLen)
	for i := range generator {
		generator[i] = letters[rand.Intn(len(letters))]
	}
	return string(generator)
}

func AuteticationUser(w http.ResponseWriter, r *http.Request) {
	user_email := r.FormValue("inputLoginEmail1")
	user_passw := r.FormValue("inputLoginPassword")
	is_auth := models.AuthenticationUser(user_email, user_passw)

	if is_auth == true {
		http.Redirect(w, r, "/user/home/", 301)
	}

	http.Redirect(w, r, "/login/", 301)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	session_id, err := r.Cookie(sessionName)
	if err != nil {
		log.Println(err)
	}
	session_id = &http.Cookie{
		Name:   sessionName,
		MaxAge: -1,
	}
	http.SetCookie(w, session_id)
	http.Redirect(w, r, "/home", 200)
}

func getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := store.Get(r, sessionName)
	if err != nil {
		clearSession(w, sessionName)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return session
}

func clearSession(w http.ResponseWriter, session string) {
	cookie := &http.Cookie{
		Name:   session,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
