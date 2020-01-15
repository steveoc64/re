package re

import "fyne.io/fyne/dataapi"

type Formation struct {
	Name             string
	SATargetModifier int
	SAFireModifier   int
}

func GetFormation(name string) (Formation, bool) {
	for _, v := range Formations.Data {
		if f, ok := v.(Formation); ok {
			if f.Name == name {
				return f, true
			}
		}
	}
	return Formation{"Unknown", 0, 0}, false
}

func (f Formation) String() string {
	return f.Name
}

func (f Formation) AddListener(func(i dataapi.DataItem)) int {
	return 0
}

func (f Formation) DeleteListener(int) {
}

var Formations = dataapi.NewSliceDataSource([]dataapi.DataItem{
	Formation{"Line", 0, 0},
	Formation{"Mixed", 1, -1},
	Formation{"Column", 3, -2},
	Formation{"Closed Col", 5, -2},
	Formation{"Square", 9, -4},
	Formation{"Skirmish", -10, 0},
	Formation{"OpenOrder", -6, 0},
})
