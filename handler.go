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

}