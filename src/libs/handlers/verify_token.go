package handlers

import (
	"net/http"
	"strings"
)

func VerifyToken(w http.ResponseWriter, r *http.Request) (string, bool) {

	// TODO: [` implement verification `]-{2023-09-29}
	token, valid := ValidateToken(r.Header.Get("Authorization"))

	if valid == 0 {
		return token, true
	} else {
		FailedToken(w, valid)
		return "", false
	}
}

func ValidateToken(raw string) (string, int8) {
	token := strings.Split(raw, " ")

	if len(token) < 2 {
		return "", 1
	}

	switch token[1] {
	case "":
		return "", 1
	case "invalid":
		return "", 2
	case "expired":
		return "", 3
	case "success":
		return token[1], 0
	}

	return "", 4
}

func FailedToken(w http.ResponseWriter, num int8) {
	http.Error(w, errorMessages[num], http.StatusUnauthorized)
}

var errorMessages = map[int8]string{
	1: "No token.",
	2: "Invalid token.",
	3: "Token expired.",
	4: "Fly me to the moon\nLet me play among the stars\nLet me see what spring is like\nOn a-Jupiter and Mars\n\nIn other words: hold my hand\nIn other words: baby, kiss me\n\nFill my heart with song\nAnd let me sing for ever more\nYou are all I long for\nAll I worship and adore\n\nIn other words: please, be true\nIn other words: I love you\n\nFill my heart with song\nLet me sing for ever more\nYou are all I long for\nAll I worship and adore\n\nIn other words: please, be true\nIn other words, in other words: I love you",
}
