package handlers

import "net/http"

func FailReq(w http.ResponseWriter, num int8) {
	msg := map[int8]string{
		0: "You sent a 'Bad Request', this is not an error on our end.",
		1: "Malformed request.",
		2: "No data found.",
		3: "User or Password, does not exist.",
		4: "Params cannot be extracted.",
		5: "User info cannot be pulled.",
		6: "Response payload cannot be parsed.",
	}
	codes := map[int8]int{
		0: http.StatusInternalServerError,
		1: http.StatusBadRequest,
		2: http.StatusNotFound,
		3: http.StatusNotFound,
		4: http.StatusBadRequest,
		5: http.StatusNotFound,
		6: http.StatusInternalServerError,
	}

	http.Error(w, msg[num], codes[num])
}
