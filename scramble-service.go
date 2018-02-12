package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yboyacigil/stringutil"

	"github.com/gorilla/mux"
)

// ScrambleResponse represents scramble service response
type ScrambleResponse struct {
	Text string `json:"text,omitempty"`
}

// GetScrambled accepts req param 'text' and returns scrambled version
func GetScrambled(w http.ResponseWriter, r *http.Request) {
	var scrambled = ScrambleResponse{""}
	vals := r.URL.Query()
	texts, ok := vals["text"]
	if ok {
		if len(texts) >= 1 {
			scrambled = ScrambleResponse{stringutil.Scramble(texts[0])}
		}
	}
	json.NewEncoder(w).Encode(scrambled)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/scramble", GetScrambled).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
