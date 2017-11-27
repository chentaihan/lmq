package server

import (
	"net/http"
	"encoding/json"
)

func SendHttpResponse(w http.ResponseWriter, m map[string]interface{}, httpCode int){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Access-Control-Allow-Headers","Content-Type")

	json.NewEncoder(w).Encode(m)
	w.WriteHeader(httpCode)
}
