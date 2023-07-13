package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := &APIResponse{
		Code: statusCode,
		Data: data,
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(respBytes)
}
