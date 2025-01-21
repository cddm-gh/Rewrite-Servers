package handlers

import (
	"fmt"
	"io"
	"net/http"
)

const oreServiceURL = "http://localhost:8625/api"

func GetAllActivities(w http.ResponseWriter, r *http.Request) {
	// Call ORE Service
	resp, err := http.Get(fmt.Sprintf("%s/v1/activities", oreServiceURL))

	if err != nil {
		http.Error(w, "Error calling ore service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error readin response", http.StatusInternalServerError)
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error from ore service %s", body), resp.StatusCode)
		return
	}

	// Return response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func GetActivityDetails(w http.ResponseWriter, r *http.Request) {
	// Get ID from path parameter
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Call ORE Service with the path parameter
	resp, err := http.Get(fmt.Sprintf("%s/v1/activities/%s", oreServiceURL, idStr))
	if err != nil {
		http.Error(w, "Error calling ore service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error from ore service: %s", body), resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
