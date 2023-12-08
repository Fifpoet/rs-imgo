package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rs-imgo/global"
)

func RunServer() {
	http.HandleFunc("/get-scale", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(global.GetScaleReq)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})
	http.HandleFunc("/get-img", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
