package pd

import (
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, error *ProblemDetails) {
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(error.Status)
	fmt.Fprint(w, error.Error())
}
