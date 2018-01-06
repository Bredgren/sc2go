package sc2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	sc2api "github.com/Bredgren/sc2go/sc2apiprotocol"
	"github.com/phayes/freeport"
)

// LocalMap returns a RequestCreateGame_LocalMap for SC2Map files.
func LocalMap(path string) *sc2api.RequestCreateGame_LocalMap {
	return &sc2api.RequestCreateGame_LocalMap{
		LocalMap: &sc2api.LocalMap{
			MapPath: path,
		},
	}
}

// LocalMapData returns RequestCreateGame_LocalMap for the given map data.
func LocalMapData(data []byte) *sc2api.RequestCreateGame_LocalMap {
	return &sc2api.RequestCreateGame_LocalMap{
		LocalMap: &sc2api.LocalMap{
			MapData: data,
		},
	}
}

// BattleNetMap returns RequestCreateGame_BattlenetMapName for the given map name.
func BattleNetMap(name string) *sc2api.RequestCreateGame_BattlenetMapName {
	return &sc2api.RequestCreateGame_BattlenetMapName{
		BattlenetMapName: name,
	}
}

// GetFreePort returns a free port. Crashes if there isn't one.
func GetFreePort() int32 {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatalf("finding a free port: %v", err)
	}
	return int32(port)
}

// GetFreePortSet returns a PortSet with free ports. Crashes if enough can't be found.
func GetFreePortSet() *sc2api.PortSet {
	return &sc2api.PortSet{
		GamePort: GetFreePort(),
		BasePort: GetFreePort(),
	}
}

// GetSC2Path returns the path to the SC2 executable and the directory it should be
// executed from.
func GetSC2Path() (exec, cwd string) {
	path := os.Getenv("SC2PATH")
	if path == "" {
		path = sc2Path
	}
	cwd = os.Getenv("SC2CWD")
	if cwd == "" {
		cwd = sc2Cwd
	}
	execPath := filepath.Join(getVersionDir(path), sc2Exec)
	cwd = filepath.Join(path, cwd)
	if _, err := os.Stat(execPath); os.IsNotExist(err) {
		log.Fatalln("No StarCraftII executable fount at", execPath)
	}
	return execPath, cwd
}

func getVersionDir(basePath string) string {
	versionDir := filepath.Join(basePath, "Versions")
	dirs, err := filepath.Glob(filepath.Join(versionDir, "Base*"))
	if err != nil {
		log.Fatalln(err)
	}
	if len(dirs) == 0 {
		log.Fatalln("No StarCraftII installation found at", basePath)
	}
	versionNumbers := []int{}
	for _, dir := range dirs {
		dirName := filepath.Base(dir)
		numStr := dirName[4:]
		verNum, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalln(err)
		}
		versionNumbers = append(versionNumbers, verNum)
	}
	sort.Ints(versionNumbers)
	latest := versionNumbers[len(versionNumbers)-1]
	return filepath.Join(basePath, "Versions", fmt.Sprintf("Base%d", latest))
}
