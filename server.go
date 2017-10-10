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

func Server(bindAddr string) {
	http.HandleFunc("/proxy", getProxy)
	err := http.ListenAndServe(bindAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
