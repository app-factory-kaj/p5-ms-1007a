package httpapi

import "net/http"

const maxNameLength = 40

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) > maxNameLength {
		writeError(w, http.StatusBadRequest, "name must be 40 characters or fewer")
		return
	}
	message := "Hello, World!"
	if name != "" {
		message = "Hello, " + name
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": message})
}
