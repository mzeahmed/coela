package symfony

import (
	"os"
	"os/exec"
)

// Install downloads and installs a fresh Symfony skeleton into dir by
// shelling out to `composer create-project symfony/skeleton`. Composer's
// own output is streamed directly to the terminal.
func Install(dir string) error {
	cmd := exec.Command("composer", "create-project", "symfony/skeleton", dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}