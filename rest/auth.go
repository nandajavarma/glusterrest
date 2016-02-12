package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

func LoadApps(fail bool) {
	grutil.Apps = make(map[string]string)
	read_apps(&grutil.Apps, true)
}

func Autoload() {
	LoadApps(true)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			LoadApps(false)
			log.Println("Reloaded")
		}
	}()
}
