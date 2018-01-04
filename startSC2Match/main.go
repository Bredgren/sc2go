package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Bredgren/sc2go/sc2"
	sc2api "github.com/Bredgren/sc2go/sc2apiprotocol"
)

var (
	seed     = uint32(time.Now().Unix())
	seedFlag int
	gameMap  = ""
)

var raceMap = map[string]sc2api.Race{
	"Random":  sc2api.Race_Random,
	"Protoss": sc2api.Race_Protoss,
	"Terran":  sc2api.Race_Terran,
	"Zerg":    sc2api.Race_Zerg,
}

var aiMap = map[string]sc2api.Difficulty{
	"VeryEasy":    sc2api.Difficulty_VeryEasy,
	"Easy":        sc2api.Difficulty_Easy,
	"Medium":      sc2api.Difficulty_Medium,
	"MediumHard":  sc2api.Difficulty_MediumHard,
	"Hard":        sc2api.Difficulty_Hard,
	"Harder":      sc2api.Difficulty_Harder,
	"VeryHard":    sc2api.Difficulty_VeryHard,
	"CheatVision": sc2api.Difficulty_CheatVision,
	"CheatMoney":  sc2api.Difficulty_CheatMoney,
	"CheatInsane": sc2api.Difficulty_CheatInsane,
}

var (
	humanRe = regexp.MustCompile(`Human_(.+)`)
	aiRe    = regexp.MustCompile(`AI_([^_]+)_(.+)`)
)

func init() {
	rand.Seed(int64(seed))
	flag.IntVar(&seedFlag, "seed", seedFlag, "The random seed to use. Default is random")
	flag.StringVar(&gameMap, "map", gameMap, "The game map to play on. Default is a random Battle.net map")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] Player1 [Player2] ...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, `
  Each Player can be one of the following:
    Human_<Race>
    AI_<Difficulty>_<Race>
    <path to bot executable>

  Where <Race> is one of:
    Random
    Protoss
    Terran
    Zerg

  And <Difficulty> is one of:
    VeryEasy
    Easy
    Medium
    MediumHard
    Hard
    Harder
    VeryHard
    CheatVision
    CheatMoney
    CheatInsane

  There can only be one Human player. If more than one bot is given then there can only be
  two of them and there can be no other human or bot players.

  The bot executable must support the following command line options:
  TODO: make these environment variables instead (partly so bots can take easily take their own options)
    -sc2Port        The port number used to connect to the SC2 cliet
    -sharedPort     The shared port number needed for joining the game
    -serverPortGame The server game port number needed for joining the game
    -serverPortBase The server base port number needed for joining the game
    -clientPortGame Comma-separated list of client game port numbers needed for joining the game
    -host           Comma-separated list of client base port numbers needed for joining the game
    -seed           Random seed

Options:
`)
		flag.PrintDefaults()
	}
}

func parseArgs() bool {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		return false
	}

	if seedFlag != 0 {
		seed = uint32(seedFlag)
	}

	log.Println("Random seed:", seed)

	return true
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	if !parseArgs() {
		return
	}

	playerArgs := flag.Args()
	playerSetup, humanCount, aiCount, botCount := getPlayerSetup(playerArgs)
	log.Println(humanCount, aiCount, botCount)
	if humanCount > 1 {
		log.Fatalln("Only one Human player allowed")
	}

	basePath := "C:/Program Files (x86)/StarCraft II"
	version := "Base60321"
	exe := "SC2_x64.exe"
	exePath := filepath.Join(basePath, "Versions", version, exe)
	cwd := filepath.Join(basePath, "Support64")

	cl, err := sc2.LaunchSC2(exePath, cwd, true)
	if err != nil {
		log.Fatalln("LaunchSC2:", err)
	}
	defer cl.Close()
	defer cl.Quit()
	log.Println(cl.GetStatus())

	if err = createGame(cl, playerSetup); err != nil {
		log.Println(err)
		return
	}

	if humanCount == 1 {
		var race sc2api.Race
		for _, setup := range playerSetup {
			if setup.GetType() == sc2api.PlayerType_Participant && setup.GetRace() != sc2api.Race_NoRace {
				race = setup.GetRace()
				break
			}
		}
		id, err := cl.JoinGame(&sc2api.RequestJoinGame{
			Participation: &sc2api.RequestJoinGame_Race{
				Race: race,
			},
			Options: &sc2api.InterfaceOptions{},
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Joined as player", id)
	} else if botCount == 0 {
		id, err := cl.JoinGame(&sc2api.RequestJoinGame{
			Participation: &sc2api.RequestJoinGame_ObservedPlayerId{
				ObservedPlayerId: 0,
			},
			Options: &sc2api.InterfaceOptions{},
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Joined as observer", id)
	}

	cl.WaitForClose()
	log.Println("Done.")
}

func getPlayerSetup(playerArgs []string) ([]*sc2api.PlayerSetup, int, int, int) {
	humanCount := 0
	aiCount := 0
	botCount := 0
	setup := []*sc2api.PlayerSetup{}
	for _, arg := range playerArgs {
		switch {
		case strings.HasPrefix(arg, "Human"):
			humanCount++
			m := humanRe.FindStringSubmatch(arg)
			if len(m) != 2 {
				log.Fatalln("Format 'Human_<Race>' expected")
			}
			race, ok := raceMap[m[1]]
			if !ok {
				log.Fatalf("Unknown race '%s'\n", m[1])
			}
			setup = append(setup, &sc2api.PlayerSetup{
				Type: sc2api.PlayerType_Participant,
				Race: race,
			})
		case strings.HasPrefix(arg, "AI"):
			aiCount++
			m := aiRe.FindStringSubmatch(arg)
			if len(m) != 3 {
				log.Fatalln("Format 'AI_<Difficulty>_<Race>' expected")
			}
			difficulty, ok := aiMap[m[1]]
			if !ok {
				log.Fatalf("Unknown difficulty '%s'\n", m[1])
			}
			race, ok := raceMap[m[2]]
			if !ok {
				log.Fatalf("Unknown race '%s'\n", m[2])
			}
			setup = append(setup, &sc2api.PlayerSetup{
				Type:       sc2api.PlayerType_Computer,
				Race:       race,
				Difficulty: difficulty,
			})
		}
	}

	if humanCount == 0 && botCount == 0 {
		setup = append(setup, &sc2api.PlayerSetup{
			Type: sc2api.PlayerType_Observer,
		})
	}
	return setup, humanCount, aiCount, botCount
}

func createGame(cl *sc2.Client, players []*sc2api.PlayerSetup) error {
	maps, err := cl.GetAvailableMaps()
	if err != nil {
		return err
	}
	fmt.Print("\nBattlenet Maps:\n", strings.Join(maps.GetBattlenetMapNames(), "\n"), "\n")
	fmt.Print("\nLocal Maps:\n", strings.Join(maps.GetLocalMapPaths(), "\n"), "\n")

	settings := &sc2api.RequestCreateGame{
		PlayerSetup: players,
		RandomSeed:  seed,
		Realtime:    true,
	}

	if gameMap == "" {
		gameMap = maps.GetBattlenetMapNames()[rand.Intn(len(maps.GetBattlenetMapNames()))]
		settings.Map = sc2.BattleNetMap(gameMap)
		log.Printf("Creating game on map '%s'\n", gameMap)
	} else {
		valid := false
		for _, mapName := range maps.GetBattlenetMapNames() {
			if gameMap == mapName {
				valid = true
				settings.Map = sc2.BattleNetMap(gameMap)
				break
			}
		}
		for _, mapName := range maps.GetLocalMapPaths() {
			if gameMap == mapName {
				valid = true
				settings.Map = sc2.LocalMap(gameMap)
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid map: %s", gameMap)
		}

	}

	return cl.CreateGame(settings)
}

// 	botProgNames := flag.Args()
//
// 	sharedPort := sc2.GetFreePort()
// 	serverPorts := sc2.GetFreePortSet()
// 	clientGamePorts := make([]string, 0, len(botProgNames))
// 	clientBasePorts := make([]string, 0, len(botProgNames))
// 	clientPorts := []*sc2api.PortSet{}
// 	for i := 0; i < len(botProgNames); i++ {
// 		gamePort := sc2.GetFreePort()
// 		basePort := sc2.GetFreePort()
// 		clientGamePorts = append(clientGamePorts, strconv.Itoa(int(gamePort)))
// 		clientBasePorts = append(clientBasePorts, strconv.Itoa(int(basePort)))
// 		clientPorts = append(clientPorts, &sc2api.PortSet{
// 			GamePort: gamePort,
// 			BasePort: basePort,
// 		})
// 	}
//
// 	basePath := "C:/Program Files (x86)/StarCraft II"
// 	version := "Base60321"
// 	exe := "SC2_x64.exe"
// 	exePath := filepath.Join(basePath, "Versions", version, exe)
// 	cwd := filepath.Join(basePath, "Support64")
//
// 	exit := make(chan struct{})
// 	cl, err := sc2.LaunchSC2(exePath, cwd, true, exit)
// 	if err != nil {
// 		log.Fatalln("LaunchSC2:", err)
// 	}
// 	defer cl.Close()
// 	defer func() {
// 		cl.Close()
// 		// Wait for application to close (not necessarily caused by cl.Close())
// 		<-exit
// 		log.Println("Done.")
// 	}()
// 	log.Println(cl.GetStatus())
//
// 	createGame(cl, clientPorts)
//
// 	cmds := []*exec.Cmd{}
// 	for i, botProgName := range botProgNames {
// 		args := []string{
// 			"-sharedPort", strconv.Itoa(int(sharedPort)),
// 			"-serverPortGame", strconv.Itoa(int(serverPorts.GamePort)),
// 			"-serverPortBase", strconv.Itoa(int(serverPorts.BasePort)),
// 			"-clientPortGame", strings.Join(clientGamePorts, ","),
// 			"-clientPortBase", strings.Join(clientBasePorts, ","),
// 		}
// 		// if i == 0 {
// 		// 	args = append(args, "-host")
// 		// }
// 		log.Println("Starting:", botProgName, args)
// 		cmd := exec.Command(botProgName, args...)
//
// 		stdout, err := cmd.StdoutPipe()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		stderr, err := cmd.StderrPipe()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if err := cmd.Start(); err != nil {
// 			log.Fatalf("start bot %d: %v\n", i, err)
// 		}
// 		cmds = append(cmds, cmd)
//
// 		go monitorOutput(i, stdout)
// 		go monitorOutput(i, stderr)
// 	}
//
// 	id, err := cl.JoinGame(&sc2api.RequestJoinGame{
// 		Participation: &sc2api.RequestJoinGame_ObservedPlayerId{
// 			ObservedPlayerId: 0,
// 		},
// 		Options: &sc2api.InterfaceOptions{},
// 		// SharedPort:  int32(sharedPort),
// 		// ServerPorts: &serverPorts,
// 		// ClientPorts: clientPorts,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println("Joined as observer", id)
//
// 	for i, cmd := range cmds {
// 		if err := cmd.Wait(); err != nil {
// 			log.Printf("Wait for bot %d to quit: %v\n", i, err)
// 		}
// 	}
// }
//
// func monitorOutput(bot int, pipe io.ReadCloser) {
// 	scanner := bufio.NewScanner(pipe)
// 	for scanner.Scan() {
// 		log.Printf("bot %d: %s\n", bot, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Printf("Error reading from bot %d: %v\n", bot, err)
// 	}
// }
//
// func createGame(cl *sc2.Client, clientPorts []*sc2api.PortSet) {
// 	maps, err := cl.GetAvailableMaps()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Print("\nBattlenet Maps:\n", strings.Join(maps.GetBattlenetMapNames(), "\n"), "\n")
// 	fmt.Print("\nLocal Maps:\n", strings.Join(maps.GetLocalMapPaths(), "\n"), "\n")
//
// 	mapChoice := maps.GetBattlenetMapNames()[rand.Intn(len(maps.GetBattlenetMapNames()))]
// 	log.Printf("Creating game on map '%s'\n", mapChoice)
// 	err = cl.CreateGame(&sc2api.RequestCreateGame{
// 		Map:         sc2.BattleNetMap(mapChoice),
// 		PlayerSetup: participants(len(clientPorts)),
// 		RandomSeed:  seed,
// 		Realtime:    true,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
//
// func participants(num int) []*sc2api.PlayerSetup {
// 	players := make([]*sc2api.PlayerSetup, 0, num)
// 	for i := 0; i < num; i++ {
// 		players = append(players, &sc2api.PlayerSetup{
// 			Type: sc2api.PlayerType_Participant,
// 		})
// 	}
// 	// players = append(players, &sc2api.PlayerSetup{
// 	// 	Type: sc2api.PlayerType_Observer,
// 	// })
// 	return players
// }
