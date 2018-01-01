package sc2

import (
	sc2api "github.com/Bredgren/sc2go/sc2apiprotocol"
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
