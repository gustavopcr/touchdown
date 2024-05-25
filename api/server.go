package api

import (
	"encoding/json"
	"fmt"
	_ "io"
	"net/http"

	"github.com/gustavopcr/touchdown/internal"
	"github.com/gustavopcr/touchdown/types"
)

func handleVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var s types.Score
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			panic("error reading from request body")
		}
		arr, err := s.ParseScore()
		if err != nil {
			panic("error parsing from request body")
		}

		res := types.NewCombination(internal.GetPos(arr[0]) * internal.GetPos(arr[1]))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)
	} else {
		fmt.Fprintf(w, "HTTP METHOD NOT SUPPORTED")
	}
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/verify", handleVerify)
	return mux
}

func NewServer(port uint16) error {
	internal.StartTouchdown()
	return http.ListenAndServe(fmt.Sprintf(":%d", port), newMux())
}
