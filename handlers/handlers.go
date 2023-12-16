// handlers/handlers.go
package handlers

import (
	"encoding/json"
	"net/http"
	"sortmodule/sorting"
)

// ProcessSingleHandler handles the /process-single endpoint.
func ProcessSingleHandler(w http.ResponseWriter, r *http.Request) {
	var payload sorting.RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	sortedArrays, timeNs := sorting.ProcessSequential(payload.ToSort)

	response := sorting.ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       timeNs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ProcessConcurrentHandler handles the /process-concurrent endpoint.
func ProcessConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	var payload sorting.RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	sortedArrays, timeNs := sorting.ProcessConcurrent(payload.ToSort)

	response := sorting.ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       timeNs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
