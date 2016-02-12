package glustercli

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type Brick struct {
	Name string `xml:"name" json:"name"`
	Uuid string `xml:"hostUuid" json:"id"`
}

type Option struct {
	Name  string `xml:"name" json:"name,omitempty"`
	Value string `xml:"value" json:"value,omitempty"`
}

type Transport string

type Volume struct {
	Name         string    `xml:"name" json:"name"`
	Id           string    `xml:"id" json:"id"`
	Status       string    `xml:"statusStr" json:"status"`
	Type         string    `xml:"typeStr" json:"type"`
	Bricks       []Brick   `xml:"bricks>brick" json:"bricks"`
	NumBricks    int       `xml:"brickCount" json:"num_bricks"`
	DistCount    int       `xml:"distCount" json:"dist_count"`
	ReplicaCount int       `xml:"replicaCount" json:"replica_count"`
	StripeCount  int       `xml:"stripeCount" json:"stripe_count"`
	TransportRaw Transport `xml:"transport" json:"transport"`
	Options      []Option  `xml:"options>option" json:"options"`
}

type Volumes struct {
	XMLName xml.Name `xml:"cliOutput"`
	List    []Volume `xml:"volInfo>volumes>volume"`
}

type VolListVolumes struct {
	XMLName xml.Name `xml:"cliOutput"`
	List    []string `xml:"volList>volume"`
}

func ExecuteCmd(cmd []string) error {
	cmd = append([]string{"--mode=script"}, cmd...)
	o, err := exec.Command("gluster", cmd...).CombinedOutput()
	if err != nil {
		return errors.New(strings.Trim(string(o), "\n"))
	}
	return nil
}

func ExecuteCmdXml(cmd []string) ([]byte, error) {
	cmd = append([]string{"--mode=script", "--xml"}, cmd...)
	o, err := exec.Command("gluster", cmd...).CombinedOutput()
	if err != nil {
		return []byte(""), errors.New(strings.Trim(string(o), "\n"))
	}
	return o, nil
}

type CreateOptions struct {
	Bricks            []string `json:"bricks"`
	ReplicaCount      int      `json:"replica"`
	StripeCount       int      `json:"stripe"`
	ArbiterCount      int      `json:"arbiter"`
	DisperseCount     int      `json:"disperse"`
	DisperseDataCount int      `json:"disperse-data"`
	RedundancyCount   int      `json:"disperse-redundancy"`
	Transport         string   `json:"transport"`
	AllowRootDir      bool     `json:"allow_root_dir"`
	ReuseBricks       bool     `json:"reuse_bricks"`
}

func VolumeCreate(volname string, bricks []string, options CreateOptions) error {
	// volume create <NEW-VOLNAME> [stripe <COUNT>] [replica <COUNT> [arbiter <COUNT>]]
	// [disperse [<COUNT>]] [disperse-data <COUNT>] [redundancy <COUNT>]
	// [transport <tcp|rdma|tcp,rdma>] <NEW-BRICK>?<vg_name>... [force]
	// - create a new volume of specified type with mentioned bricks
	// TODO Validate the Inputs and Options
	cmd := []string{"volume", "create", volname}
	if options.ReplicaCount != 0 {
		cmd = append(cmd, "replica", fmt.Sprintf("%d", options.ReplicaCount))
	}
	if options.StripeCount != 0 {
		cmd = append(cmd, "stripe", fmt.Sprintf("%d", options.StripeCount))
	}
	if options.ArbiterCount != 0 {
		cmd = append(cmd, "arbiter", fmt.Sprintf("%d", options.ArbiterCount))
	}
	if options.DisperseCount != 0 {
		cmd = append(cmd, "disperse", fmt.Sprintf("%d", options.DisperseCount))
	}
	if options.DisperseDataCount != 0 {
		cmd = append(cmd, "disperse-data", fmt.Sprintf("%d", options.DisperseDataCount))
	}
	if options.RedundancyCount != 0 {
		cmd = append(cmd, "redundancy", fmt.Sprintf("%d", options.RedundancyCount))
	}
	if options.Transport != "" {
		cmd = append(cmd, "transport", options.Transport)
	}

	cmd = append(cmd, bricks...)

	if options.AllowRootDir || options.ReuseBricks {
		// Limitation in Gluster 1.0, Can't differenciate between
		// multiple Flags. Common option "force"
		cmd = append(cmd, "force")
	}

	return ExecuteCmd(cmd)
}

func VolumeStart(volname string) error {
	cmd := []string{"volume", "start", volname}
	return ExecuteCmd(cmd)
}

func VolumeStop(volname string) error {
	cmd := []string{"volume", "stop", volname}
	return ExecuteCmd(cmd)
}

func VolumeDelete(volname string) error {
	cmd := []string{"volume", "delete", volname}
	return ExecuteCmd(cmd)
}

func VolumeOptSet(volname string, key string, value string) error {
	cmd := []string{"volume", "set", volname, key, value}
	return ExecuteCmd(cmd)
}

func VolumeOptReset(volname string, key string, force bool) error {
	cmd := []string{"volume", "reset", volname}
	if key != "" {
		cmd = append(cmd, key)
	}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

func VolumeOptGet() {

}

func VolumeLogRotate(volname string, brick string) error {
	cmd := []string{"volume", "log", volname, "rotate"}
	if brick != "" {
		cmd = append(cmd, brick)
	}
	return ExecuteCmd(cmd)
}

func VolumeRestart(volname string) error {
	err_stop := VolumeStop(volname)
	if err_stop != nil {
		return err_stop
	}
	return VolumeStart(volname)
}

func VolumeInfo(volname string) ([]Volume, error) {
	var q Volumes
	cmd := []string{"volume", "info"}
	if volname != "" {
		cmd = append(cmd, volname)
	}
	data, err := ExecuteCmdXml(cmd)
	if err != nil {
		return []Volume{}, err
	}
	xmlerr := xml.Unmarshal(data, &q)
	if xmlerr != nil {
		return []Volume{}, xmlerr
	}
	return q.List, nil
}

func VolumeStatus() {
	// Merge Vol Status and Vol Info to get Offline Status
}

func VolumeList() ([]string, error) {
	var q VolListVolumes
	cmd := []string{"volume", "list"}
	data, err := ExecuteCmdXml(cmd)
	if err != nil {
		return []string{}, err
	}
	xmlerr := xml.Unmarshal(data, &q)
	if xmlerr != nil {
		return []string{}, xmlerr
	}
	return q.List, nil
}

func VolumeClearLocks() {

}

func VolumeBarrierEnable(volname string) error {
	cmd := []string{"volume", "barrier", volname, "enable"}
	return ExecuteCmd(cmd)
}

func VolumeBarrierDisable(volname string) error {
	cmd := []string{"volume", "barrier", volname, "disable"}
	return ExecuteCmd(cmd)
}
