// volume bitrot <VOLNAME> {enable|disable} |
// volume bitrot <volname> scrub-throttle {lazy|normal|aggressive} |
// volume bitrot <volname> scrub-frequency {hourly|daily|weekly|biweekly|monthly} |
// volume bitrot <volname> scrub {pause|resume|status} - Bitrot translator specific operation. For more information about bitrot command type  'man gluster'

package glustercli

var (
	ThrottleTypeLazy       = "lazy"
	ThrottleTypeNormal     = "normal"
	ThrottleTypeAggressive = "aggressive"
	ScrubFrequencyHourly   = "hourly"
	ScrubFrequencyDaily    = "daily"
	ScrubFrequencyWeekly   = "weekly"
	ScrubFrequencyBiweekly = "biweekly"
	ScrubFrequencyMonthly  = "monthly"
)

func VolumeBitrotEnable(volname string) error {
	cmd := []string{"volume", "bitrot", volname, "enable"}
	return ExecuteCmd(cmd)
}

func VolumeBitrotDisable(volname string) error {
	cmd := []string{"volume", "bitrot", volname, "disable"}
	return ExecuteCmd(cmd)
}

func VolumeBitrotScrubThrottle(volname string, throttle_type string) error {
	cmd := []string{"volume", "bitrot", volname, "scrub-throttle", throttle_type}
	return ExecuteCmd(cmd)
}

func VolumeBitrotScrubThrottleLazy(volname string) error {
	return VolumeBitrotScrubThrottle(volname, ThrottleTypeLazy)
}

func VolumeBitrotScrubThrottleNormal(volname string) error {
	return VolumeBitrotScrubThrottle(volname, ThrottleTypeNormal)
}

func VolumeBitrotScrubThrottleAggressive(volname string) error {
	return VolumeBitrotScrubThrottle(volname, ThrottleTypeAggressive)
}

func VolumeBitrotScrubFrequency(volname string, frequency_type string) error {
	cmd := []string{"volume", "bitrot", volname, "scrub-frequency", frequency_type}
	return ExecuteCmd(cmd)
}

func VolumeBitrotScrubFrequencyHourly(volname string) error {
	return VolumeBitrotScrubFrequency(volname, ScrubFrequencyDaily)
}

func VolumeBitrotScrubFrequencyDaily(volname string) error {
	return VolumeBitrotScrubFrequency(volname, ScrubFrequencyDaily)
}

func VolumeBitrotScrubFrequencyWeekly(volname string) error {
	return VolumeBitrotScrubFrequency(volname, ScrubFrequencyWeekly)
}

func VolumeBitrotScrubFrequencyBiweekly(volname string) error {
	return VolumeBitrotScrubFrequency(volname, ScrubFrequencyBiweekly)
}

func VolumeBitrotScrubFrequencyMonthly(volname string) error {
	return VolumeBitrotScrubFrequency(volname, ScrubFrequencyMonthly)
}

func VolumeBitrotScrubPause(volname string) error {
	cmd := []string{"volume", "bitrot", volname, "scrub", "pause"}
	return ExecuteCmd(cmd)
}

func VolumeBitrotScrubResume(volname string) error {
	cmd := []string{"volume", "bitrot", volname, "scrub", "resume"}
	return ExecuteCmd(cmd)
}

func VolumeBitrotScrubStatus() {

}
