package handlers

import "net/http"

func FailReq(w http.ResponseWriter, num int8) {
	msg := map[int8]string{
		0: "You sent a 'Bad Request', this is not an error on our end.",
		1: "Malformed request.",
		2: "No data found.",
		3: "User or Password, does not exist.",
	}

	http.Error(w, msg[num], http.StatusBadRequest)
}
