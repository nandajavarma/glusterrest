package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/aravindavk/glusterrest/grutil"
)

func write_config(config *grutil.Config, fail bool) {
	lock, err := grutil.NewLock(grutil.CONF_FILE)
	if err != nil {
		if fail {
			log.Fatal("Config file Lock failed", err)
		}
		return
	}
	defer lock.Unlock()
	lock.Lock()
	data_op, err2 := json.Marshal(config)
	if err2 != nil {
		if fail {
			log.Fatal("to json err", err2)
		}
		return
	}
	err3 := ioutil.WriteFile(grutil.CONF_FILE, data_op, 0644)
	if err3 != nil {
		if fail {
			log.Fatal("Write err", err3)
		}
	}
}

func read_config(config *grutil.Config, fail bool) {
	lock, err := grutil.NewLock(grutil.CONF_FILE)
	if err != nil {
		if fail {
			log.Fatal("Config file Lock failed", err)
		}
		return
	}
	defer lock.Unlock()
	lock.Lock()
	data, err := ioutil.ReadFile(grutil.CONF_FILE)
	if err != nil && fail {
		log.Fatal("No conf file")
	}

	err1 := json.Unmarshal(data, &config)
	if err1 != nil && fail {
		log.Fatal("json err")
	}
}
