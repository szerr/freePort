package proxy

import (
	"encoding/json"
	"log"
	"net/http"
)

func getProxy(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(GetAllProxy())
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

func Server() {
	http.HandleFunc("/proxy", getProxy)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
