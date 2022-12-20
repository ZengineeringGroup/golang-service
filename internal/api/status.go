package api

import (
	"fmt"
	"net/http"
)

func (a *API) StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/status")
	w.WriteHeader(http.StatusOK)
}
