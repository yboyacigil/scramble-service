package scramble

import (
	"encoding/json"
	"net/http"

	"github.com/yboyacigil/stringutil"
)

// Response represents scramble service response
type Response struct {
	Text string `json:"text,omitempty"`
}

// GetScrambled accepts req param 'text' and returns scrambled version
func GetScrambled(w http.ResponseWriter, r *http.Request) {
	var scrambled = Response{""}
	vals := r.URL.Query()
	texts, ok := vals["text"]
	if ok {
		if len(texts) >= 1 {
			scrambled = Response{stringutil.Scramble(texts[0])}
		}
	}
	json.NewEncoder(w).Encode(scrambled)
}

func init() {
	http.HandleFunc("/", GetScrambled)
}
