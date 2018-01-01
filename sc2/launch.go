package sc2

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strconv"
	"time"

	"github.com/phayes/freeport"
	"golang.org/x/net/websocket"
)

const timeout = 120

// LaunchSC2 starts the SC2 client and returns a connection to it for communicating.
// exePath should be the full path to the SC2 executable. cwd is a path to the directory
// to run SC2 from. If cwd is the empty string then the current directory is used. If
// windowed is false then the SC2 client will start in fullscreen mode. When the SC2 client
// exits one value will be sent over the exit channel.
func LaunchSC2(exePath, cwd string, windowed bool, exit chan<- struct{}) (*websocket.Conn, error) {
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

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", freePort), Path: "/sc2api"}

	origin := "http://localhost/"
	var conn *websocket.Conn
	for secondsToTry := timeout; secondsToTry > 0; secondsToTry-- {
		conn, err = websocket.Dial(u.String(), "", origin)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if conn == nil {
		return nil, fmt.Errorf("timed out connecting to SC2")
	}

	go func() {
		if err := cmd.Wait(); err != nil {
			log.Println("Wait for SC2 client to quit:", err)
		}
		exit <- struct{}{}
	}()

	return conn, nil
}
