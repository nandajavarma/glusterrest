package glustercli

// volume tier <VOLNAME> attach [<replica COUNT>] <NEW-BRICK>...
// volume tier <VOLNAME> detach <start|stop|status|commit|[force]>
//  - Tier translator specific operations.
// volume attach-tier <VOLNAME> [<replica COUNT>] <NEW-BRICK>... - NOTE: this is old syntax, will be depreciated in next release. Please use gluster volume tier <vol> attach [<replica COUNT>] <
// NEW-BRICK>...
// volume detach-tier <VOLNAME>  <start|stop|status|commit|force> - NOTE: this is old syntax, will be depreciated in next release. Please use gluster volume tier <vol> detach {start|stop|commit
// } [force]
// volume add-brick <VOLNAME> [<stripe|replica> <COUNT>] <NEW-BRICK> ... [force] - add brick to volume <VOLNAME>
// volume remove-brick <VOLNAME> [replica <COUNT>] <BRICK> ... <start|stop|status|commit|force> - remove brick from volume <VOLNAME>
// volume rebalance <VOLNAME> {{fix-layout start} | {start [force]|stop|status}} - rebalance operations
// volume replace-brick <VOLNAME> <SOURCE-BRICK> <NEW-BRICK> {commit force} - replace-brick operations

func VolumeAddBrick(volname string, bricks []string, replica_count int, stripe_count int) error {
	cmd := []string{"volume", "add-brick", volname, "scrub", "pause"}
	return ExecuteCmd(cmd)
}

func VolumeRemoveBrickStart() {

}

func VolumeRemoveBrickStop() {

}

func VolumeRemoveBrickCommit() {

}

func VolumeRemoveBrickStatus() {

}

func VolumeRebalanceFixLayoutStart(volname string) error {
	cmd := []string{"volume", "rebalance", volname, "fix-layout", "start"}
	return ExecuteCmd(cmd)
}

func VolumeRebalanceStart(volname string) error {
	cmd := []string{"volume", "rebalance", volname, "start"}
	return ExecuteCmd(cmd)
}

func VolumeRebalanceStop(volname string) error {
	cmd := []string{"volume", "rebalance", volname, "stop"}
	return ExecuteCmd(cmd)
}

func VolumeRebalanceStatus() {

}

func VolumeReplaceBrick() {

}

func VolumeTierAttach() {

}

func VolumeTierDitachStart() {

}

func VolumeTierDitachStop() {

}

func VolumeTierDitachStatus() {

}

func VolumeTierDitachCommit() {

}
