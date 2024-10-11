package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var pongNumber = 0

// func writeNumberToFile(filepath string, number int) error {
// 	// Convert the number to a string
// 	data := fmt.Sprintf("%d\n", number)
// 	err := os.WriteFile(filepath, []byte(data), 0644)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

type pongResponse struct {
	Number int `json:"number"`
}

func main() {
	r := mux.NewRouter()
	// url := os.Getenv("FILE_URL")
	pong := pongResponse{
		Number: pongNumber,
	}

	r.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "pong %v", pongNumber)

		// err := writeNumberToFile(url, pongNumber)
		// if err != nil {
		// 	fmt.Printf("Error writing to file: %v\n", err)
		// } else {
		// 	fmt.Printf("Number %d written to file %s successfully\n", pongNumber, url)
		// }
		responseWithJSON(w, http.StatusOK, pong)
		pongNumber++
		pong.Number = pongNumber
	}).Methods("GET")

	r.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		responseWithJSON(w, http.StatusOK, pong)
	}).Methods("GET")

	http.ListenAndServe(":3000", r)
} // TODO: Write the internal services and test it
