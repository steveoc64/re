package re

import (
	"fyne.io/fyne/dataapi"
)

type Terrain struct {
	Name             string
	SAFireModifier   int
	SATargetModifier int
}

func (t Terrain) String() string {
	return t.Name
}

func (t Terrain) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (t Terrain) DeleteListener(int) {
}

func GetTerrain(str string) (Terrain, bool) {
	for _, v := range Terrains.Data {
		if t, ok := v.(Terrain); ok {
			if t.Name == str {
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
	Terrain{"Clear", 0, 1},
	Terrain{"LightWoods", -1, -2},
	Terrain{"Woods", -2, -2},
	Terrain{"HeavyWoods", -3, -2},
	Terrain{"LightCover", -1, -4},
	Terrain{"Cover", -1, -8},
	Terrain{"HeavyCover", 0, -12},
	Terrain{"SuperHeavyCover", 1, -16},
})
