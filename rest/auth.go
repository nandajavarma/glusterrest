package main

import (
	"encoding/json"
	"net/http"

	"github.com/aravindavk/glusterrest/grutil"
)

func HttpErrorJSON(w http.ResponseWriter, err string, code int) {
	msg := grutil.ErrorResponse{Message: err}
	j, _ := json.Marshal(msg)
	w.WriteHeader(code)
	w.Write(j)
}

func ValidateSign(secret string, message string, sign string) bool {
	app_sign := grutil.Sign(secret, message)
	return (sign == app_sign)
}
