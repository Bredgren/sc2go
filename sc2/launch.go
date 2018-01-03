package sc2

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/phayes/freeport"
)

const timeout = 120

// LaunchSC2 starts the SC2 client and returns a Client object for communicating. exePath
// should be the full path to the SC2 executable. cwd is a path to the directory to run
// SC2 from. If cwd is the empty string then the current directory is used. If windowed
// is false then the SC2 client will start in fullscreen mode. When the SC2 client exits
// for any reason one value will be sent over the exit channel.
func LaunchSC2(exePath, cwd string, windowed bool) (*Client, error) {
	freePort, err := freeport.GetFreePort()
	if err != nil {
		return nil, fmt.Errorf("finding a free port: %v", err)
	}

	mode := "1"
	if windowed {
		mode = "0"
	}

	args := []string{
		"-listen", "127.0.0.1",
		"-port", strconv.Itoa(freePort),
		"-displayMode", mode,
	}
	cmd := exec.Command(exePath, args...)
	if cwd != "" {
		cmd.Dir = cwd
	}
	if err = cmd.Start(); err != nil {
		return nil, fmt.Errorf("start SC2: %v", err)
	}

	return NewClient(freePort)
}
