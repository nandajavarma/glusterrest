package grutil

import (
	"github.com/aravindavk/glusterrest/glustercli"
	"github.com/gorilla/websocket"
)

var (
	PORT          = 8080
	CERT          = "/var/lib/glusterd/rest/server.pem"
	KEY           = "/var/lib/glusterd/rest/server.key"
	HTTPS         = false
	APPS_DB       = "/var/lib/glusterd/rest/apps.json"
	CONF_FILE     = "/etc/glusterfs/glusterrest.json"
	SERVER_LOG    = "/var/log/glusterfs/rest/access.log"
	EVENTS_SOCK   = "/var/run/gluster/events.sock"
	INTERNAL_USER = "gluster"
	INTERNAL_URL  = "/v1/listen"
	PEERS_LIST    = "peers.json"

	EventMessages = make(chan string, 10)
	Peers         []glustercli.Peer

	VolumeStatusCreate = "Created"
	VolumeStatusStart  = "Started"
	VolumeStatusStop   = "Stopped"
	VolumeStatusDelete = ""
	NODE_ID            = "N1"
	EventVolumeCreate  = "Volume.Create"
	EventVolumeStart   = "Volume.Start"
	EventVolumeStop    = "Volume.Stop"
	EventVolumeDelete  = "Volume.Delete"

	Apps map[string]string

	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	WS_clients = WsClients{}
)
