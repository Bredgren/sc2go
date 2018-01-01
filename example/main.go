package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/Bredgren/sc2go/sc2"
)

func main() {
	basePath := "C:/Program Files (x86)/StarCraft II"
	version := "Base60321"
	exe := "SC2_x64.exe"
	exePath := filepath.Join(basePath, "Versions", version, exe)
	cwd := filepath.Join(basePath, "Support64")

	exit := make(chan struct{})

	cl, err := sc2.LaunchSC2(exePath, cwd, true, exit)
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

	cl.Quit()

	<-exit

	log.Println("Done.")
}
