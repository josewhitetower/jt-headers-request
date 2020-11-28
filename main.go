package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
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

	port := getPort()

	log.Println("Server running in port: " + port)
	log.Fatal(http.ListenAndServe(port, nil))

}

// GetPort the Port from the environment so we can run on Heroku
func getPort() string {
	port := os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
