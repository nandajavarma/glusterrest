package main

// No Previous State check if it is User driven event

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/aravindavk/glusterrest/glustercli"
	"github.com/aravindavk/glusterrest/grutil"
)

func Sign(secret string, message string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func HandleVolumeCreate(data string) {
	fmt.Printf("Volume %s is Created\n", data)
	grutil.SendEvent("Volume.Create=" + data)
}

func HandleVolumeStart(data string) {
	fmt.Printf("Volume %s is Started\n", data)
}

func HandleVolumeStop(data string) {
	fmt.Printf("Volume %s is Stopped\n", data)
}

func HandleVolumeDelete(data string) {
	fmt.Printf("Volume %s is Deleted\n", data)
}

func LoadPeers(fail bool) {
	read_peers(true)
}

func Autoload() {
	LoadPeers(true)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			LoadPeers(false)
			log.Println("Reloaded")
		}
	}()
}

func read_peers(fail bool) {
	p, err := glustercli.PoolList()
	if err != nil {
		if fail {
			log.Fatal("Peers list failed", err)
		}
	}
	grutil.Peers = p
}

func worker() {
	for {
		parse_line(<-grutil.EventMessages)
	}
}

func reader(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		grutil.EventMessages <- string(buf[0:nr])
		_, err = c.Write([]byte("0"))
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func parse_line(line string) {
	msg_parts := strings.Split(line, "=")
	switch msg_parts[0] {
	case grutil.EventVolumeCreate:
		HandleVolumeCreate(msg_parts[1])
	case grutil.EventVolumeStart:
		HandleVolumeStart(msg_parts[1])
	case grutil.EventVolumeStop:
		HandleVolumeStop(msg_parts[1])
	case grutil.EventVolumeDelete:
		HandleVolumeDelete(msg_parts[1])
	}
}

func events_listener(server_address string) {
	cleanup_err := os.Remove(server_address)
	if cleanup_err != nil {
		fmt.Println(cleanup_err)
	}
	l, err := net.Listen("unix", server_address)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go worker()

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go reader(fd)
	}
}

func read_gluster_app_id(apps *map[string]string, fail bool) {
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

func main() {
	Autoload()
	events_listener(grutil.EVENTS_SOCK)
}
