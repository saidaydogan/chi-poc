package controller

import "net/http"

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetById"))
}

func GetDetailById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetDetailById"))
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateById"))
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteById"))
}
