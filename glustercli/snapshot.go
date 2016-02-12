// snapshot create <snapname> <volname> [no-timestamp] [description <description>] [force] - Snapshot Create.
// snapshot clone <clonename> <snapname> - Snapshot Clone.
// snapshot restore <snapname> - Snapshot Restore.
// snapshot status [(snapname | volume <volname>)] - Snapshot Status.
// snapshot info [(snapname | volume <volname>)] - Snapshot Info.
// snapshot list [volname] - Snapshot List.
// snapshot config [volname] ([snap-max-hard-limit <count>] [snap-max-soft-limit <percent>]) | ([auto-delete <enable|disable>])| ([activate-on-create <enable|disable>]) - Snapshot Config.
// snapshot delete (all | snapname | volume <volname>) - Snapshot Delete.
// snapshot activate <snapname> [force] - Activate snapshot volume.
// snapshot deactivate <snapname> - Deactivate snapshot volume.

package glustercli

func VolumeSnapshotCreate(snapname string, volname string) {

}

func SnapshotConfigMaxHardLimit(volname string, count int) error {
	cmd := []string{"snapshot", "config", volname, "snap-max-hard-limit", string(count)}
	return ExecuteCmd(cmd)
}

func SnapshotConfigMaxSoftLimit(volname string, percent int) error {
	cmd := []string{"snapshot", "config", volname, "snap-max-soft-limit", string(percent)}
	return ExecuteCmd(cmd)
}

func SnapshotConfigAutoDeleteEnable(volname string) error {
	cmd := []string{"snapshot", "config", volname, "auto-delete", "enable"}
	return ExecuteCmd(cmd)
}

func SnapshotConfigAutoDeleteDisable(volname string) error {
	cmd := []string{"snapshot", "config", volname, "auto-delete", "disable"}
	return ExecuteCmd(cmd)
}

func SnapshotConfigActivateOnCreateEnable(volname string) error {
	cmd := []string{"snapshot", "config", volname, "activate-on-create", "enable"}
	return ExecuteCmd(cmd)
}

func SnapshotConfigActivateOnCreateDisable(volname string) error {
	cmd := []string{"snapshot", "config", volname, "activate-on-create", "disable"}
	return ExecuteCmd(cmd)
}

func VolumeSnapshotDeleteAll() error {
	cmd := []string{"snapshot", "delete", "all"}
	return ExecuteCmd(cmd)
}

func VolumeSnapshotDeleteBySnapname(snapname string) error {
	cmd := []string{"snapshot", "delete", snapname}
	return ExecuteCmd(cmd)
}

func VolumeSnapshotDeleteByVolume(volname string) error {
	cmd := []string{"snapshot", "delete", volname}
	return ExecuteCmd(cmd)
}

func VolumeSnapshotClone(snapname string, clonename string) error {
	cmd := []string{"snapshot", "clone", clonename, snapname}
	return ExecuteCmd(cmd)
}

func VolumeSnapshotRestore(snapname string) error {
	cmd := []string{"snapshot", "restore", snapname}
	return ExecuteCmd(cmd)
}

func VolumeSnapshotStatus() {

}

func VolumeSnapshotInfo() {

}

func VolumeSnapshotList() {

}

func VolumeSnapshotActivate(snapname string, force bool) error {
	cmd := []string{"snapshot", "activate", snapname}
	if force {
		cmd = append(cmd, "force")
	}
	return ExecuteCmd(cmd)
}

func VolumeSnapshotDeactivate(snapname string) error {
	cmd := []string{"snapshot", "deactivate", snapname}
	return ExecuteCmd(cmd)
}
