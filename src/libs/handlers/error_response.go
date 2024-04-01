package handlers

import (
	"net/http"
)

type CustomError struct {
	Code    int
	Message string
}

func FailReq(args ...any) {
	w := args[0].(http.ResponseWriter)
	num := args[1].(int)

	msg := map[int]string{
		0:  "You sent a 'Bad Request'",
		1:  "Malformed request.",
		2:  "No data found.",
		3:  "Data, does not exist.",
		4:  "Params cannot be extracted.",
		5:  "User info cannot be pulled.",
		6:  "Response payload cannot be parsed.",
		7:  "Unknown url",
		99: "Error not possible",
	}
	codes := map[int]int{
		0:   http.StatusInternalServerError,
		1:   http.StatusBadRequest,
		2:   http.StatusNotFound,
		3:   http.StatusNotFound,
		4:   http.StatusNotFound,
		5:   http.StatusNotFound,
		6:   http.StatusInternalServerError,
		7:   http.StatusInternalServerError,
		99:  http.StatusNoContent,
		101: http.StatusNotAcceptable,
	}

	var message string

	if num >= 101 {
		str, _ := JSON(CustomError{Code: num, Message: args[2].(error).Error()})
		message = str
	} else {
		str, _ := JSON(CustomError{Code: num, Message: msg[num]})
		message = str
	}

	http.Error(w, message, codes[num])
}
