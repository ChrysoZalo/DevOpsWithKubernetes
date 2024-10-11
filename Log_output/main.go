package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var strSlice = []string{}

func randomString(length int) string {
	ran_str := make([]byte, length)

	// Generate a random string consisting of uppercase letters (A-Z)
	for i := 0; i < length; i++ {
		// Generate a random ASCII value for uppercase letters A-Z (65 to 90)
		ran_str[i] = byte(65 + r.Intn(26))
	}
	return string(ran_str)
}

func logging() []string {
	length := 24
	ran_str := randomString(length)
	for {
		strSlice = append(strSlice, ran_str)
		time.Sleep(5 * time.Second)
	}
}

type pongResponse struct {
	Number int `json:"number"`
}

func main() {
	r := mux.NewRouter()
	// url := os.Getenv("FILE_URL")

	go logging()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, str := range strSlice {
			fmt.Fprintln(w, str)
			data, err := http.Get("http://localhost:3000/pong")
			if err != nil || data.Status == "404 Not Found" {
				responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving data from pingpong api:", err))
			}
			// data, err := os.ReadFile(url)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			decoder := json.NewDecoder(data.Body)
			pongResponse := pongResponse{}
			err = decoder.Decode(&pongResponse)
			if err != nil {
				responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error parsing JSON:", err))
			}
			fmt.Fprintln(w, pongResponse.Number)
		}
	}).Methods("GET")

	http.ListenAndServe(":80", r)
}
