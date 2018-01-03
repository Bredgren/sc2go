// A simple bot that works with startSC2Match.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Bredgren/sc2go/sc2"
	sc2api "github.com/Bredgren/sc2go/sc2apiprotocol"
)

var (
	seed           = uint32(time.Now().Unix())
	seedFlag       int
	sharedPort     int
	serverPortGame int
	serverPortBase int
	clientPortGame string
	clientPortBase string
	host           = false
)

var (
	serverPorts sc2api.PortSet
	clientPorts []*sc2api.PortSet
)

func init() {
	rand.Seed(int64(seed))
	flag.IntVar(&seedFlag, "seed", seedFlag, "")
	flag.IntVar(&sharedPort, "sharedPort", sharedPort, "")
	flag.IntVar(&serverPortGame, "serverPortGame", serverPortGame, "")
	flag.IntVar(&serverPortBase, "serverPortBase", serverPortBase, "")
	flag.StringVar(&clientPortGame, "clientPortGame", clientPortGame, "")
	flag.StringVar(&clientPortBase, "clientPortBase", clientPortBase, "")
	flag.BoolVar(&host, "host", host, "Start as game host")
}

func parseFlags() {
	flag.Parse()

	if seedFlag != 0 {
		seed = uint32(seedFlag)
	}

	serverPorts.GamePort = int32(serverPortGame)
	serverPorts.BasePort = int32(serverPortBase)

	clientPortGameStrs := strings.Split(clientPortGame, ",")
	clientPortBaseStrs := strings.Split(clientPortBase, ",")
	if len(clientPortGameStrs) != len(clientPortBaseStrs) {
		log.Fatal("Client ports lists not the same size")
	}
	for i := 0; i < len(clientPortGameStrs); i++ {
		gamePort, err := strconv.Atoi(clientPortGameStrs[i])
		if err != nil {
			log.Fatal(err)
		}
		basePort, err := strconv.Atoi(clientPortBaseStrs[i])
		if err != nil {
			log.Fatal(err)
		}
		clientPorts = append(clientPorts, &sc2api.PortSet{
			GamePort: int32(gamePort),
			BasePort: int32(basePort),
		})
	}
}

func main() {
	parseFlags()
	fmt.Println("Random Seed:", seed)

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
	log.Println(cl.GetStatus())

	if host {
		createGame(cl)
	}

	id, err := cl.JoinGame(&sc2api.RequestJoinGame{
		Participation: &sc2api.RequestJoinGame_Race{
			Race: sc2api.Race_Random,
		},
		Options:     &sc2api.InterfaceOptions{},
		SharedPort:  int32(sharedPort),
		ServerPorts: &serverPorts,
		ClientPorts: clientPorts,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Joined as player", id)

	cl.WaitForClose()
	log.Println("Done.")
}

func createGame(cl *sc2.Client) {
	maps, err := cl.GetAvailableMaps()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nBattlenet Maps:\n", strings.Join(maps.GetBattlenetMapNames(), "\n"), "\n")
	fmt.Print("\nLocal Maps:\n", strings.Join(maps.GetLocalMapPaths(), "\n"), "\n")

	mapChoice := maps.GetBattlenetMapNames()[rand.Intn(len(maps.GetBattlenetMapNames()))]
	log.Printf("Creating game on map '%s'\n", mapChoice)
	err = cl.CreateGame(&sc2api.RequestCreateGame{
		Map:         sc2.BattleNetMap(mapChoice),
		PlayerSetup: participants(len(clientPorts)),
		RandomSeed:  seed,
		Realtime:    true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func participants(num int) []*sc2api.PlayerSetup {
	players := make([]*sc2api.PlayerSetup, 0, num)
	for i := 0; i < num; i++ {
		players = append(players, &sc2api.PlayerSetup{
			Type: sc2api.PlayerType_Participant,
		})
	}
	return players
}
