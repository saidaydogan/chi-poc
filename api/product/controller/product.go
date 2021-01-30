package controller

import "net/http"

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetById"))
}

func GetDetailById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetDetailById"))
}
