package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/aravindavk/glusterrest/grutil"
)

func read_apps(apps *map[string]string, fail bool) {
	lock, err := grutil.NewLock(grutil.APPS_DB)
	if err != nil {
		if fail {
			log.Fatal("Apps file Lock failed", err)
		}
		return
	}
	defer lock.Unlock()
	lock.Lock()
	data, err := ioutil.ReadFile(grutil.APPS_DB)
	if err != nil && fail {
		log.Fatal("No file")
	}

	err1 := json.Unmarshal(data, &apps)
	if err1 != nil && fail {
		log.Fatal("json err")
	}
}

func write_apps(apps *map[string]string, fail bool) {
	lock, err := grutil.NewLock(grutil.APPS_DB)
	if err != nil {
		if fail {
			log.Fatal("Apps file Lock failed", err)
		}
		return
	}
	defer lock.Unlock()
	lock.Lock()
	data_op, err2 := json.Marshal(apps)
	if err2 != nil {
		if fail {
			log.Fatal("to json err", err2)
		}
		return
	}
	err3 := ioutil.WriteFile(grutil.APPS_DB, data_op, 0644)
	if err3 != nil {
		if fail {
			log.Fatal("Write err", err3)
		}
	}
}
