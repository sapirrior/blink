package engine

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func normalizeURL(target string) string {
	if strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "http://") {
		return target
	}

	return "https://" + target
}

func isWSL() bool {
	return os.Getenv("WSL_DISTRO_NAME") != ""
}

func Open(target string) error {
	var cmd *exec.Cmd
	target = normalizeURL(target)

	switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", target)
		case "linux":
				if isWSL() {
					cmd = exec.Command("explorer.exe", target)
				} else {
					cmd = exec.Command("xdg-open", target)
				}
		case "darwin":
			cmd = exec.Command("open", target)
	}

	return cmd.Start()
}