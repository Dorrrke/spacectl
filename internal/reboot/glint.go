package reboot

import "os/exec"

func RebootGlintLauncher() error {
	cmd := exec.Command("systemctl", "restart", "glint-launcher.service")
	err := cmd.Run()
	return err
}
