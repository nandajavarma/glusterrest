package glustercli

// volume quota <VOLNAME> {enable|disable|list [<path> ...]| list-objects [<path> ...] | remove <path>| remove-objects <path> | default-soft-limit <percent>} |
// volume quota <VOLNAME> {limit-usage <path> <size> [<percent>]} |
// volume quota <VOLNAME> {limit-objects <path> <number> [<percent>]} |
// volume quota <VOLNAME> {alert-time|soft-timeout|hard-timeout} {<time>} - quota translator specific operations
// volume inode-quota <VOLNAME> enable - quota translator specific operations

func VolumeQuotaEnable(volname string) error {
	cmd := []string{"volume", "quota", volname, "enable"}
	return ExecuteCmd(cmd)
}

func VolumeQuotaDisable(volname string) error {
	cmd := []string{"volume", "quota", volname, "disable"}
	return ExecuteCmd(cmd)
}

func VolumeQuotaList() {

}

func VolumeQuotaLimitUsage() {

}

func VolumeQuotaLimitObjects() {

}

func VolumeQuotaAlertTime() {

}

func VolumeQuotaSoftTimeout() {

}

func VolumeQuotaHardTimeout() {

}

func VolumeInodeQuotaEnable() {

}

func VolumeInodeQuotaDisable() {

}
