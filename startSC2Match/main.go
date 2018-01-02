package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Bredgren/sc2go/sc2"
)

// var (
// 	seed = 0
// )

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] BOT1 BOT2 ...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, `
This program starts multiple SC2 bot programs passing them data that need they need
to share.

There must be at least 2 BOTS. All BOTS must be paths to executables that support
the following command line options:
	-sharedPort      A number
	-serverPortGame  A number
	-serverPortBase  A number
	-clientPortGame  Comma-separated list of numbers
	-clientPortBase  Comma-separated list of numbers
	-host            If present then the bot should be host
	-seed            Random seed

The first BOT will be made host.
`)
		flag.PrintDefaults()
	}
}

func parseArgs() bool {
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		return false
	}

	return true
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	if !parseArgs() {
		return
	}

	botProgNames := flag.Args()

	sharedPort := sc2.GetFreePort()
	serverPorts := sc2.GetFreePortSet()
	clientGamePorts := make([]string, 0, len(botProgNames))
	clientBasePorts := make([]string, 0, len(botProgNames))
	for i := 0; i < len(botProgNames); i++ {
		clientGamePorts = append(clientGamePorts, strconv.Itoa(int(sc2.GetFreePort())))
		clientBasePorts = append(clientBasePorts, strconv.Itoa(int(sc2.GetFreePort())))
	}

	cmds := []*exec.Cmd{}
	for i, botProgName := range botProgNames {
		args := []string{
			"-sharedPort", strconv.Itoa(int(sharedPort)),
			"-serverPortGame", strconv.Itoa(int(serverPorts.GamePort)),
			"-serverPortBase", strconv.Itoa(int(serverPorts.BasePort)),
			"-clientPortGame", strings.Join(clientGamePorts, ","),
			"-clientPortBase", strings.Join(clientBasePorts, ","),
		}
		if i == 0 {
			args = append(args, "-host")
		}
		log.Println("Starting:", botProgName, args)
		cmd := exec.Command(botProgName, args...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatalf("start bot %d: %v\n", i, err)
		}
		cmds = append(cmds, cmd)

		go monitorOutput(i, stdout)
		go monitorOutput(i, stderr)
	}

	for i, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			log.Printf("Wait for bot %d to quit: %v\n", i, err)
		}
	}
}

func monitorOutput(bot int, pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		log.Printf("bot %d: %s\n", bot, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from bot %d: %v\n", bot, err)
	}
}
