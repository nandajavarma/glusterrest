// volume geo-replication [<VOLNAME>] [<SLAVE-URL>] {create [[ssh-port n] [[no-verify]|[push-pem]]] [force]|start [force]|stop [force]|pause [force]|resume [force]|config|status [detail]|delete} [options...] - Geo-sync operations

package glustercli

func GsecCreate() error {
	cmd := []string{"system::", "execute", "gsec-create"}
	return ExecuteCmd(cmd)
}

func VolumeGeorepCreatePushPem(mastervol string, slave string, slavevol string, force bool) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "create", "push-pem"}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

func VolumeGeorepCreateNoVerify(mastervol string, slave string, slavevol string, force bool) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "create", "no-verify"}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

func VolumeGeorepStart(mastervol string, slave string, slavevol string, force bool) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "start"}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

func VolumeGeorepStop(mastervol string, slave string, slavevol string, force bool) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "stop"}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)

}

func VolumeGeorepPause(mastervol string, slave string, slavevol string, force bool) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "pause"}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

func VolumeGeorepResume(mastervol string, slave string, slavevol string, force bool) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "resume"}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

func VolumeGeorepDelete(mastervol string, slave string, slavevol string, force bool) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "delete"}
	return ExecuteCmd(cmd)
}

func VolumeGeorepStatus() {
	// Merge Geo-rep Status + VolInfo
}

func VolumeGeorepConfigGet() {

}

func VolumeGeorepConfigSet(mastervol string, slave string, slavevol string, key string, value string) error {
	cmd := []string{"volume", "geo-replication", mastervol, slave + "::" + slavevol, "config", key, value}
	return ExecuteCmd(cmd)
}
