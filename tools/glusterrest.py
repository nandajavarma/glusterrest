#!/usr/bin/env python

# package main

# import (
# 	"encoding/json"
# 	"fmt"
# 	"log"
# 	"net/http"
# 	"os"

# 	"github.com/gorilla/mux"
# )

# func usage_add() {
# 	fmt.Println("USAGE: glusterrest add NAME [SECRET]")
# }

# func handle_add() {
# 	// Check for existence in file
# 	// If exists ERROR
# 	// If not exists Add to Records and Generate a Secret Key and PRINT
# 	if len(os.Args) < 3 {
# 		log.Fatal("Invalid App Name")
# 	}
# 	appname := os.Args[2]
# 	var token string
# 	if len(os.Args) == 4 {
# 		token = os.Args[3]
# 	} else {
# 		token = randToken()
# 	}
# 	var apps map[string]string
# 	read_apps(&apps, true)

# 	if _, ok := apps[appname]; !ok {
# 		apps[appname] = token
# 		write_apps(&apps, true)
# 		fmt.Println("Your app secret is", token)
# 	} else {
# 		fmt.Println("App Exists")
# 	}
# }

# func handle_delete() {
# 	// Check for existence in Records
# 	// If exists DELETE
# 	// If not exists ERROR
# 	if len(os.Args) < 3 {
# 		log.Fatal("Invalid App Name")
# 	}
# 	appname := os.Args[2]
# 	var apps map[string]string
# 	read_apps(&apps, true)

# 	if _, ok := apps[appname]; ok {
# 		delete(apps, appname)
# 		write_apps(&apps, true)
# 	} else {
# 		fmt.Println("App not Exists")
# 	}
# }

# func handle_reset() {
# 	// Check for existence in Records
# 	// If not exists ERROR
# 	// If exists generate new Secret Key and update Records
# 	if len(os.Args) < 3 {
# 		log.Fatal("Invalid App Name")
# 	}
# 	appname := os.Args[2]
# 	var apps map[string]string
# 	read_apps(&apps, true)

# 	var token string
# 	if len(os.Args) == 4 {
# 		token = os.Args[3]
# 	} else {
# 		token = randToken()
# 	}

# 	if _, ok := apps[appname]; ok {
# 		apps[appname] = token
# 		write_apps(&apps, true)
# 		fmt.Println("Your new app secret is", token)
# 	} else {
# 		fmt.Println("App not Exists")
# 	}
# }

# func handle_config_set(name string, value []byte) {
# 	var config Config
# 	read_config(&config, true)

# 	if name == "port" {
# 		json.Unmarshal(value, &config.Port)
# 	} else if name == "https" {
# 		json.Unmarshal(value, &config.Https)
# 	}
# 	write_config(&config, true)
# }

# func handle_config_get(name string) {
# 	var config Config
# 	read_config(&config, true)
# 	if name == "" || name == "port" {
# 		fmt.Println("Port: ", config.Port)
# 	}
# 	if name == "" || name == "https" {
# 		fmt.Println("HTTPS: ", config.Https)
# 	}
# }

# func handle_config() {
# 	name := ""
# 	if len(os.Args) >= 3 {
# 		name = os.Args[2]
# 	}

# 	if len(os.Args) == 4 {
# 		handle_config_set(name, []byte(os.Args[3]))
# 	} else {
# 		handle_config_get(name)
# 	}
# }

# func handle_server() {
# 	Autoload()
# 	router := mux.NewRouter().StrictSlash(true)
# 	AddRoutes(router)

# 	port_data := fmt.Sprintf(":%d", PORT)
# 	if HTTPS {
# 		log.Fatal(http.ListenAndServeTLS(port_data, CERT, KEY, nil))
# 	} else {
# 		log.Fatal(http.ListenAndServe(port_data, nil))
# 	}
# }

# func handle_cert_gen() {

# }
