package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Fenix/internal/models"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const sessionName = "sessionid"

var CookieLen = 32

func AuteticationUser(w http.ResponseWriter, r *http.Request) {
	user_email := r.FormValue("inputLoginEmail")
	user_passw := r.FormValue("inputLoginPassword")

	is_auth := models.AuthenticationUser(user_email, user_passw)

	if is_auth == true {
		newSession(w, r, user_email)
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


func GetSession(w http.ResponseWriter, r *http.Request) *http.Cookie {
	session, err := r.Cookie(sessionName)
	fmt.Println(session)
	fmt.Println(err)

	if err != nil {
		clearSession(w, sessionName)
		http.Redirect(w, r, "/home/", http.StatusTemporaryRedirect)
	}

	return session
}

func newSession(w http.ResponseWriter, r *http.Request, email string) {
	store := sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
	store.Options.Path = "/user/home/"
	store.New(r, sessionName)
	session, err := store.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		clearSession(w, sessionName)
		http.Redirect(w, r, "/home/", http.StatusTemporaryRedirect)
	}

	session.Values["useremail"] = email
	session.Save(r, w)
	store.Save(r, w, session)

	expires := time.Now().Add(3600)
	
	userid := models.SelectUserId(email)
	models.InsertAuth(userid, session.ID, expires)
}

func clearSession(w http.ResponseWriter, session string) {
	cookie := &http.Cookie{
		Name:   session,
		Value:  "",
		Path: "/home/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func newSession2(w http.ResponseWriter, r *http.Request) {
	newCookie := &http.Cookie{
		Name: sessionName,
		Value: cookieGenerator(),
		MaxAge: 3600,
	}
	r.AddCookie(newCookie)
	http.SetCookie(w, newCookie)
}

func cookieGenerator() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	generator := make([]byte, CookieLen)
	for i := range generator {
		generator[i] = letters[rand.Intn(len(letters))]
	}
	return string(generator)
}