package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/whoami", func(w http.ResponseWriter, r *http.Request) {

		type Response struct {
			Language  string `json:"language"`
			Software  string `json:"software"`
			IPAddress string `json:"ipaddress"`
		}

		res := Response{r.Header.Values("Accept-Language")[0], r.Header.Values("User-Agent")[0], r.RemoteAddr}
		js, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	})

	const port = 8080

	log.Println("Server running in port " + fmt.Sprint(port))
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(port), nil))

}
