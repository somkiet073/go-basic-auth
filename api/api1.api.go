package api

import (
	"fmt"
	"net/http"
)

func Demo1API(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Demo 1 API")
}
