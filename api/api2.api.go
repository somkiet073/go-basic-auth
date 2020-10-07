package api

import (
	"fmt"
	"net/http"
)

func Demo2API(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Demo 2 API")
}
