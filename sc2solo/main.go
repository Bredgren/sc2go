package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Bredgren/sc2go/sc2"
	sc2api "github.com/Bredgren/sc2go/sc2apiprotocol"
)

var seed uint32

var (
	gameMap      string
	race         = "random"
	aiRace       = ""
	aiDifficulty = "veryeasy"
)

var raceMap = map[string]sc2api.Race{
	"random":  sc2api.Race_Random,
	"protoss": sc2api.Race_Protoss,
	"terran":  sc2api.Race_Terran,
	"zerg":    sc2api.Race_Zerg,
}

var difficultyMap = map[string]sc2api.Difficulty{
	"veryeasy":    sc2api.Difficulty_VeryEasy,
	"easy":        sc2api.Difficulty_Easy,
	"medium":      sc2api.Difficulty_Medium,
	"mediumhard":  sc2api.Difficulty_MediumHard,
	"hard":        sc2api.Difficulty_Hard,
	"harder":      sc2api.Difficulty_Harder,
	"veryhard":    sc2api.Difficulty_VeryHard,
	"cheatvision": sc2api.Difficulty_CheatVision,
	"cheatmoney":  sc2api.Difficulty_CheatMoney,
	"cheatinsane": sc2api.Difficulty_CheatInsane,
}

func init() {
	flag.StringVar(&gameMap, "map", gameMap, "Map (default <random Battle.net map>)")
	flag.StringVar(&race, "race", race, "Your race")
	flag.StringVar(&aiRace, "ai", aiRace, "AI race (default <no AI opponent>)")
	flag.StringVar(&aiDifficulty, "d", aiDifficulty, "AI difficulty")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, `
Play SC2 solo or against an AI.

Race options (case insensitive):
  Random
  Protoss
  Terran
  Zerg

AI difficulty options (case insensitive):
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

Options:
`)
		flag.PrintDefaults()
	}

	seed = uint32(time.Now().Unix())
	rand.Seed(int64(seed))
}

func parseFlags() {
	flag.Parse()

	race = strings.ToLower(race)
	aiRace = strings.ToLower(aiRace)
	aiDifficulty = strings.ToLower(aiDifficulty)
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	parseFlags()

	path, cwd, err := sc2.GetSC2Path()
	if err != nil {
		log.Fatalln(err)
	}
	cl, err := sc2.LaunchSC2(path, cwd, false)
	if err != nil {
		log.Fatalln("Connect to client:", err)
	}
	defer cl.Close()
	defer cl.Quit()

	joinSettings := &sc2api.RequestJoinGame{
		Participation: &sc2api.RequestJoinGame_Race{
			Race: raceMap[race],
		},
		Options: &sc2api.InterfaceOptions{},
	}

	players := []*sc2api.PlayerSetup{
		{
			Type: sc2api.PlayerType_Participant,
		},
	}
	if aiRace != "" {
		players = append(players, &sc2api.PlayerSetup{
			Type:       sc2api.PlayerType_Computer,
			Race:       raceMap[aiRace],
			Difficulty: difficultyMap[aiDifficulty],
		})
	}

	settings := &sc2api.RequestCreateGame{
		PlayerSetup: players,
		RandomSeed:  seed,
		Realtime:    true,
	}
	sc2.SetMap(cl, settings, gameMap)

	log.Println("\nStarting game")
	log.Println("Press Alt+F4 to quit")
	err = cl.CreateGame(settings)
	if err != nil {
		log.Fatalln("Create game:", err)
	}

	_, err = cl.JoinGame(joinSettings)
	if err != nil {
		log.Fatal("Join game:", err)
	}

	cl.WaitForClose()
	log.Println("Done.")
}
