package handlers

import (
	"kaab/src/libs/fixtures"
	"kaab/src/models"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.ContentLength == 0 {
			deleteCookie(w)
			FailReq(w, 1)
			return
		}

		var data models.LogInPayload
		err := GetBody(r, &data)
		if err != nil {
			deleteCookie(w)
			FailReq(w, 2)
		}

		if checkUserLogin(data) {

			token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6Im1lY3VhdGUiLCJpYXQiOjE1MTYyMzkwMjJ9.ShuL8v0cyrBurYVVYrAqpMYKkK0m33nTZ5oULVPY-Uw"
			makeCookie(w, token)
			w.Write([]byte(token))
			return
		} else {
			deleteCookie(w)
			FailReq(w, 3)
			return
		}
	}
	deleteCookie(w)
	FailReq(w, 0)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	deleteCookie(w)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func checkUserLogin(data models.LogInPayload) bool {
	pool := fixtures.UserData

	for i := 0; i < len(pool); i++ {
		if pool[i].Email == data.Email && pool[i].Password == data.Password {
			return true
		}
	}

	return false
}

func makeCookie(w http.ResponseWriter, token string) {
	auth_cookie := &http.Cookie{
		Name:  "session",
		Value: token,
		// Path       string    // optional
		// Domain     string    // optional
		// Expires    time.Time // optional
		// RawExpires string    // for reading cookies only
		MaxAge:   1800,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Raw:      token,
	}
	http.SetCookie(w, auth_cookie)
}

func deleteCookie(w http.ResponseWriter) {
	clear_cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Raw:      "",
	}
	http.SetCookie(w, clear_cookie)
}
