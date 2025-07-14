package reboot

import "os/exec"

func RebootSpaceAgent() error {
	cmd := exec.Command("systemctl", "restart", "space-agent-pc.service")
	err := cmd.Run()
	return err
}
