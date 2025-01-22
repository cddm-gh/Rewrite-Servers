package handlers

import (
	"fmt"
	"io"
	"net/http"
	"orem/config"
)

func GetAllActivities(w http.ResponseWriter, r *http.Request) {
	// Create a new request to forward
	url := fmt.Sprintf("%s/v1/resort/activities?%s", config.Get().OREServiceURL, r.URL.RawQuery)
	req, err := http.NewRequestWithContext(r.Context(), "GET", url, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Forward all headers from the original request
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error calling ore service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
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

	// Create a new request to forward
	url := fmt.Sprintf("%s/v1/activities/%s%s", config.Get().OREServiceURL, idStr, r.URL.RawQuery)
	req, err := http.NewRequestWithContext(r.Context(), "GET", url, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Forward all headers from the original request
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Make the request
	resp, err := http.DefaultClient.Do(req)
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
