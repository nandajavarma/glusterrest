// peer probe { <HOSTNAME> | <IP-address> } - probe peer specified by <HOSTNAME>
// peer detach { <HOSTNAME> | <IP-address> } [force] - detach peer specified by <HOSTNAME>
// peer status - list status of peers
// peer help - Help command for peer
// pool list - list all the nodes in the pool (including localhost)

package glustercli

import (
	"encoding/xml"
)

func PeerAttach(host string) error {
	cmd := []string{"peer", "probe", host}
	return ExecuteCmd(cmd)
}

func PeerDitach(host string) error {
	cmd := []string{"peer", "detach", host}
	return ExecuteCmd(cmd)
}

type Peer struct {
	Id        string `xml:"uuid" json:"id"`
	Hostname  string `xml:"hostname" json:"hostname"`
	Connected int    `xml:"connected" json:"connected"`
}

type Peers struct {
	XMLName xml.Name `xml:"cliOutput"`
	List    []Peer   `xml:"peerStatus>peer"`
}

func PeerStatus() {

}

func PoolList() ([]Peer, error) {
	var q Peers
	cmd := []string{"pool", "list"}

	data, err := ExecuteCmdXml(cmd)
	if err != nil {
		return []Peer{}, err
	}
	xmlerr := xml.Unmarshal(data, &q)
	if xmlerr != nil {
		return []Peer{}, xmlerr
	}
	return q.List, nil
}
