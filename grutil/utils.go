package grutil

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"log"
	"os/signal"
	"io/ioutil"

	"github.com/gorilla/websocket"
	"github.com/aravindavk/glusterrest/glustercli"
)

func read_apps(apps *map[string]string, fail bool) {
	lock, err := NewLock(APPS_DB)
	if err != nil {
		if fail {
			log.Fatal("Apps file Lock failed", err)
		}
		return
	}
	defer lock.Unlock()
	lock.Lock()
	data, err := ioutil.ReadFile(APPS_DB)
	if err != nil && fail {
		log.Fatal("No file")
	}

	err1 := json.Unmarshal(data, &apps)
	if err1 != nil && fail {
		log.Fatal("json err")
	}
}

func write_apps(apps *map[string]string, fail bool) {
	lock, err := NewLock(APPS_DB)
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
	err3 := ioutil.WriteFile(APPS_DB, data_op, 0644)
	if err3 != nil {
		if fail {
			log.Fatal("Write err", err3)
		}
	}
}

func LoadApps(fail bool) {
	Apps = make(map[string]string)
	read_apps(&Apps, true)
}

func Autoload() {
	LoadApps(true)
	LoadPeers(true)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			LoadApps(false)
			LoadPeers(true)
			log.Println("Reloaded")
		}
	}()
}

func LoadPeers(fail bool) {
	read_peers(true)
}

func read_peers(fail bool) {
	p, err := glustercli.PoolList()
	if err != nil {
		if fail {
			log.Fatal("Peers list failed", err)
		}
	}
	Peers = p
}

func write_config(config *Config, fail bool) {
	lock, err := NewLock(CONF_FILE)
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
	err3 := ioutil.WriteFile(CONF_FILE, data_op, 0644)
	if err3 != nil {
		if fail {
			log.Fatal("Write err", err3)
		}
	}
}

func read_config(config *Config, fail bool) {
	lock, err := NewLock(CONF_FILE)
	if err != nil {
		if fail {
			log.Fatal("Config file Lock failed", err)
		}
		return
	}
	defer lock.Unlock()
	lock.Lock()
	data, err := ioutil.ReadFile(CONF_FILE)
	if err != nil && fail {
		log.Fatal("No conf file")
	}

	err1 := json.Unmarshal(data, &config)
	if err1 != nil && fail {
		log.Fatal("json err")
	}
}

func NewLock(path string) (*Lockfile, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &Lockfile{path: path, fh: fh}, nil
}

func (lock *Lockfile) Lock() error {
	lock.Mutex.Lock()
	defer lock.Mutex.Unlock()
	if lock.fh == nil {
		var err error
		if lock.fh, err = os.Open(lock.path); err != nil {
			return err
		}
	}
	err := syscall.Flock(int(lock.fh.Fd()), syscall.LOCK_EX)
	return err
}

func (lock *Lockfile) Unlock() error {
	lock.Mutex.Lock()
	defer lock.Mutex.Unlock()
	if lock.fh == nil {
		return nil
	}
	err := syscall.Flock(int(lock.fh.Fd()), syscall.LOCK_UN)
	lock.fh.Close()
	lock.fh = nil
	return err
}

func (c *WsClients) Add(conn *websocket.Conn) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.clients = append(c.clients, conn)
}

func (c *WsClients) Remove(idx int) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.clients = append(c.clients[:idx], c.clients[idx+1:]...)
}

func (c *WsClients) SendAll(msg []byte) {
	c.Mutex.Lock()
	clients := c.clients
	c.Mutex.Unlock()
	for idx, ws_client := range clients {
		err := ws_client.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			c.Remove(idx)
		}
	}
}

func Execute(cmd []string) Resp {
	cmd = append([]string{"--mode=script"}, cmd...)
	out := Resp{Ok: true}
	o, err := exec.Command("gluster", cmd...).CombinedOutput()
	if err != nil {
		out.Ok = false
		out.Msg = strings.Trim(string(o), "\n")
	}
	return out
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func Sign(secret string, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func http_post(host string, message string, ch chan *HttpResponse, client *http.Client) {
	fmt.Println(host)
	api_url := "/v1/listen"
	url := "http://" + host + ":8080" + api_url
	dte := time.Now().UTC().String()
	string_to_sign := "POST\napplication/json\n" + dte + "\n" + api_url
	data := EventMsg{Node: NODE_ID, Message: message}
	data_json, err := json.Marshal(data)
	if err != nil {
		return
	}
	sign := Sign(Apps["gluster"], string_to_sign)
	r, _ := http.NewRequest("POST", url, bytes.NewBufferString(string(data_json)))
	r.Header.Add("Authorization", "HMAC_SHA256 gluster:"+sign)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Date", dte)
	resp, err := client.Do(r)
	ch <- &HttpResponse{url, resp, err}
	if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
		resp.Body.Close()
	}
}

func send_event_all(message string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	client := http.Client{}
	for _, url := range Peers {
		go http_post(url.Hostname, message, ch, &client)
	}
	if len(Peers) == 0 {
		return responses
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			if r.err != nil {
				fmt.Println("with an error", r.err)
			}
			responses = append(responses, r)
			if len(responses) == len(Peers) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func SendEvent(message string) {
	result := send_event_all(message)
	for _, result := range result {
		if result != nil && result.response != nil {
			fmt.Printf("%s status: %s\n", result.url,
				result.response.Status)
		}
	}
}
