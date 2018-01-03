package main

import (
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	"github.com/Bredgren/sc2go/sc2"
	sc2api "github.com/Bredgren/sc2go/sc2apiprotocol"
)

var seed = uint32(time.Now().Unix())

func init() {
	rand.Seed(int64(seed))
}

func main() {
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

	maps, err := cl.GetAvailableMaps()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nBattlenet Maps:\n", strings.Join(maps.GetBattlenetMapNames(), "\n"), "\n")
	fmt.Print("\nLocal Maps:\n", strings.Join(maps.GetLocalMapPaths(), "\n"), "\n")

	ping, err := cl.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Base Build", ping.GetBaseBuild())
	log.Println("Data Build", ping.GetDataBuild())
	log.Println("Game Version", ping.GetGameVersion())
	log.Println("Data Version", ping.GetDataVersion())

	mapChoice := maps.GetBattlenetMapNames()[rand.Intn(len(maps.GetBattlenetMapNames()))]
	log.Printf("Creating game on map '%s'\n", mapChoice)
	err = cl.CreateGame(&sc2api.RequestCreateGame{
		Map: sc2.BattleNetMap(mapChoice),
		PlayerSetup: []*sc2api.PlayerSetup{
			// {
			// 	Type: sc2api.PlayerType_Participant,
			// },
			{
				Type: sc2api.PlayerType_Observer,
			},
			{
				Type:       sc2api.PlayerType_Computer,
				Race:       sc2api.Race_Random,
				Difficulty: sc2api.Difficulty_CheatInsane,
			},
			{
				Type:       sc2api.PlayerType_Computer,
				Race:       sc2api.Race_Random,
				Difficulty: sc2api.Difficulty_VeryEasy,
			},
		},
		RandomSeed: seed,
		Realtime:   true,
	})
	if err != nil {
		log.Fatal(err)
	}

	settings := &sc2api.RequestJoinGame{
		// Participation: &sc2api.RequestJoinGame_Race{
		// 	Race: sc2api.Race_Random,
		// },
		Participation: &sc2api.RequestJoinGame_ObservedPlayerId{
			ObservedPlayerId: 0,
		},
		Options: &sc2api.InterfaceOptions{},
	}
	id, err := cl.JoinGame(settings)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Joined as player", id)

	cl.WaitForClose()
	log.Println("Done.")
}
