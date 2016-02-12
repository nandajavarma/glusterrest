// nfs-ganesha {enable| disable}  - Enable/disable NFS-Ganesha support

package glustercli

func NfsGaneshaEnable() error {
	cmd := []string{"nfs-ganesha", "enable"}
	return ExecuteCmd(cmd)
}

func NfsGaneshaDisable() error {
	cmd := []string{"nfs-ganesha", "disable"}
	return ExecuteCmd(cmd)
}
