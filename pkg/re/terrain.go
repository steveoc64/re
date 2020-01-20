package re

import (
	"fmt"
	"fyne.io/fyne/dataapi"
)

type Terrain struct {
	Name             string
	SAFireModifier   int
	SATargetModifier int
	MoraleMod int
}

func (t Terrain) String() string {
	return t.Name
}

func (t Terrain) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (t Terrain) DeleteListener(int) {
}

func GetTerrain(str fmt.Stringer) (Terrain, bool) {
	for _, v := range Terrains.Data {
		if t, ok := v.(Terrain); ok {
			if t.Name == str.String() {
				return t, true
			}
		}
	}
	return Terrain{
		Name:             "Unknown",
		SAFireModifier:   0,
		SATargetModifier: 0,
	}, false
}

var Terrains = dataapi.NewSliceDataSource([]dataapi.DataItem{
	Terrain{"Open", 0, 1, 0},
	Terrain{"LightWoods", -1, -2, 0},
	Terrain{"Woods", -2, -4, -1},
	Terrain{"HeavyWoods", -3, -8, -2},
	Terrain{"LightCover", -1, -4, 1},
	Terrain{"Cover", -1, -8, 2},
	Terrain{"HeavyCover", 0, -12, 3},
	Terrain{"SuperHeavyCover", 1, -16, 4},
})
