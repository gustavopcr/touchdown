package api

import (
	"fmt"
	"net/http"

	"github.com/gustavopcr/touchdown/internal"
)

func handleVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, fmt.Sprintf("%d", internal.GetPos(3)))
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
