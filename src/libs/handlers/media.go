package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func StaticPublicAssetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unknown request :: StaticMediaAssetHandler %s", r.Header)
}

func StaticMediaHandler(r *mux.Router, path string) {
	redirPath := DirectoryTable["STATIC_PUBLIC_MEDIA"]
	r.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir(redirPath))))
}

func StaticVideoHandler(r *mux.Router, path string) {
	redirPath := "assets/video/"
	r.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir(redirPath))))
}

func StaticFormattedMedia(r *mux.Router, path string) {
	r.PathPrefix(path).Handler(http.StripPrefix(path, http.HandlerFunc(ResizeFormatHandler)))
}

func ResizeFormatHandler(w http.ResponseWriter, r *http.Request) {

	req_assets := strings.Split(r.URL.String(), "/")

	if len(req_assets) == 3 && validateMediaFormattingURLString(req_assets[0]) {
		redirPath := DirectoryTable["STATIC_FORMATTED_MEDIA"]
		mydir, _ := os.Getwd()
		img_path := fmt.Sprintf("%s/%s%s", mydir, redirPath, req_assets[2])

		pic_values := strings.Split(req_assets[0], "x")
		x, _ := strconv.Atoi(pic_values[0])
		y, _ := strconv.Atoi(pic_values[1])
		dat := ResImg(img_path, x, y)

		w.WriteHeader(http.StatusOK)
		w.Write(dat)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
