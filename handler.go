package main

import (
"encoding/json"
"net/http"
"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func Handler(w http.ResponseWriter, req *http.Request){
	currentTime := time.Now()
	response := TimeResponse(Time: currentTime.Format(time.RFC3339))
	jsonResponse, err := json.MarshalIndent(response, "", " ")

	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return;
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}