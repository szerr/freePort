package proxy

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetProxy(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(GetAllProxy())
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

func DelProxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	DeleteProxy(r.FormValue("key"))
}

func Server() {
	http.HandleFunc("/getproxy", GetProxy)
	http.HandleFunc("/delproxy", DelProxy)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
