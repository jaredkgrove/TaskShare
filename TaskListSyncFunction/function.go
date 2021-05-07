// Package p contains an HTTP Cloud Function. Some new stuff
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "Hello bird!")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	if d.Message == "" {
		fmt.Fprint(w, "Hello bird!")
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
