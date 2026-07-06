package browser

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func isWSL() bool {
	return os.Getenv("WSL_DISTRO_NAME") != ""
}

// Open launches the default system browser to navigate to the target URL or path.
func Open(target string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", target)
	case "darwin":
		cmd = exec.Command("open", target)
	case "linux":
		if isWSL() {
			cmd = exec.Command("explorer.exe", target)
		} else {
			cmd = exec.Command("xdg-open", target)
		}
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	return cmd.Start()
}
