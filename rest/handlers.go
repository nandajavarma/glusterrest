package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"

	"github.com/aravindavk/glusterrest/glustercli"
	"github.com/aravindavk/glusterrest/grutil"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SignHandler(h http.Handler) http.Handler {
	// Middleware to handle HMAC sign Verification on all requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the Signature here, This will be applied to All Routes
		authHeader := strings.TrimSpace(r.Header.Get("Authorization"))
		if authHeader == "" {
			HttpErrorJSON(w, "Missing 'Authorization' header", 401)
			return
		}
		auth_data := strings.Split(authHeader, " ")
		app_id_sign := strings.Split(auth_data[1], ":")

		// When App ID/Name not present in Apps list - Unauthorized
		if _, ok := grutil.Apps[app_id_sign[0]]; !ok {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Special Case for Internal APIs Only AppId:gluster can send message
		if app_id_sign[0] != grutil.INTERNAL_USER && r.URL.Path == grutil.INTERNAL_URL {
			http.Error(w, http.StatusText(403), 403)
			return
		}

		// TODO: Compare Date header with Timeout and raise 403 if timeout

		var buffer bytes.Buffer

		buffer.WriteString(r.Method)
		buffer.WriteString("\n")

		buffer.WriteString(r.Header.Get("Content-Type"))
		buffer.WriteString("\n")

		buffer.WriteString(r.Header.Get("Date"))
		buffer.WriteString("\n")

		buffer.WriteString(r.URL.Path)
		// buffer.WriteString("\n")

		// If Valid App ID but Signature not Matching then Forbidden
		if !ValidateSign(grutil.Apps[app_id_sign[0]], buffer.String(), app_id_sign[1]) {
			HttpErrorJSON(w, http.StatusText(403), 403)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func RestLoggingHandler(h http.Handler) http.Handler {
	// Middleware to log the incoming request as per the Apache Common logging framework
	logFile, err := os.OpenFile(grutil.SERVER_LOG, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, h)
}

func SetApplicationHeaderJson(h http.Handler) http.Handler {
	// Middleware to set Application Header as JSON
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func ListenHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var msg grutil.EventMsg
	err := decoder.Decode(&msg)
	if err != nil {
		HttpErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	grutil.WS_clients.SendAll([]byte(msg.Message))
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grutil.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		//log.Println(err)
		return
	}

	grutil.WS_clients.Add(conn)
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return
		}
	}
}

func VolumeCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var opts glustercli.CreateOptions
	err := decoder.Decode(&opts)
	if err != nil {
		HttpErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	volName := vars["volName"]
	err_create := glustercli.VolumeCreate(volName, opts.Bricks, opts)
	if err_create != nil {
		HttpErrorJSON(w, err_create.Error(), http.StatusInternalServerError)
		return
	}
}

func VolumeGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName, ok := vars["volName"]
	if ok {
		w.Write([]byte("With Volume Name" + volName))
	} else {
		w.Write([]byte("Without Volume Name"))
	}
}

func VolumeStart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName := vars["volName"]
	err := glustercli.VolumeStart(volName)
	if err != nil {
		HttpErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func VolumeStop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName := vars["volName"]
	err := glustercli.VolumeStop(volName)
	if err != nil {
		HttpErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func VolumeDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volName := vars["volName"]
	err := glustercli.VolumeDelete(volName)
	if err != nil {
		HttpErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
