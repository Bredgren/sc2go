package sc2

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

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
// executed from. Checks environment variables SC2PATH and SC2CWD before trying platform
// defaults.
func GetSC2Path() (exec, cwd string, err error) {
	path := os.Getenv("SC2PATH")
	if path == "" {
		path = sc2Path
	}
	cwd = os.Getenv("SC2CWD")
	if cwd == "" {
		cwd = sc2Cwd
	}
	verDir, err := getVersionDir(path)
	if err != nil {
		return "", "", err
	}
	execPath := filepath.Join(verDir, sc2Exec)
	cwd = filepath.Join(path, cwd)
	if _, err := os.Stat(execPath); os.IsNotExist(err) {
		return "", "", fmt.Errorf("no StarCraftII executable fount at %s", execPath)
	}
	return execPath, cwd, nil
}

func getVersionDir(basePath string) (string, error) {
	versionDir := filepath.Join(basePath, "Versions")
	dirs, err := filepath.Glob(filepath.Join(versionDir, "Base*"))
	if err != nil {
		return "", err
	}
	if len(dirs) == 0 {
		return "", fmt.Errorf("no StarCraftII installation found at %s", basePath)
	}
	versionNumbers := []int{}
	for _, dir := range dirs {
		dirName := filepath.Base(dir)
		numStr := dirName[4:]
		verNum, err := strconv.Atoi(numStr)
		if err != nil {
			return "", err
		}
		versionNumbers = append(versionNumbers, verNum)
	}
	sort.Ints(versionNumbers)
	latest := versionNumbers[len(versionNumbers)-1]
	return filepath.Join(basePath, "Versions", fmt.Sprintf("Base%d", latest)), nil
}

// SetMap prints out available maps then if m is empty it sets settings.Map to a random
// Battle.net map, and if m is not empty it sets it to the map named by m (local or Battle.net)
func SetMap(cl *Client, settings *sc2api.RequestCreateGame, m string) {
	validMaps, err := cl.GetAvailableMaps()
	if err != nil {
		log.Fatalln("Get available maps:", err)
	}
	fmt.Print("\nBattlenet Maps:\n", strings.Join(validMaps.GetBattlenetMapNames(), "\n"), "\n")
	fmt.Print("\nLocal Maps:\n", strings.Join(validMaps.GetLocalMapPaths(), "\n"), "\n")

	if m == "" {
		settings.Map = BattleNetMap(validMaps.GetBattlenetMapNames()[rand.Intn(len(validMaps.GetBattlenetMapNames()))])
	} else {
		for _, mapName := range validMaps.GetBattlenetMapNames() {
			if m == mapName {
				settings.Map = BattleNetMap(m)
				break
			}
		}
		for _, mapName := range validMaps.GetLocalMapPaths() {
			if m == mapName {
				settings.Map = LocalMap(m)
				break
			}
		}
	}
}
