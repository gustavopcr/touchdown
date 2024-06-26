package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavopcr/touchdown/internal"
	"github.com/gustavopcr/touchdown/types"
)

func HandleVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var s types.Score
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, "error reading from request body", http.StatusBadRequest)
			return
		}
		arr, err := internal.ParseScore(s.Score)
		if err != nil {
			http.Error(w, "error parsing from request body", http.StatusBadRequest)
			return
		}
		res := types.NewCombination(internal.GetPos(arr[0]) * internal.GetPos(arr[1]))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)
	} else {
		fmt.Fprintf(w, "HTTP METHOD NOT SUPPORTED")
	}
}
