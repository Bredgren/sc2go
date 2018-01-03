package sc2

import (
	"log"

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
func GetFreePortSet() sc2api.PortSet {
	return sc2api.PortSet{
		GamePort: GetFreePort(),
		BasePort: GetFreePort(),
	}
}
