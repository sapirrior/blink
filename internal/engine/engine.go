package engine

import (
	"os"
	"os/exec"
	"runtime"
)


func isWSL() bool {
	return os.Getenv("WSL_DISTRO_NAME") != ""
}

func Open(target string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("runlld32", "url.dll,FileProtocolHandler", target)
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