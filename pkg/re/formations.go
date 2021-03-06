package re

import (
	"fmt"
	"fyne.io/fyne/dataapi"
)

type Formation struct {
	Name             string
	SATargetModifier int
	SAFireModifier   int
	MaxFire          string
}

func GetFormation(name fmt.Stringer) (Formation, bool) {
	for _, v := range Formations.Data {
		if f, ok := v.(Formation); ok {
			if f.Name == name.String() {
				return f, true
			}
		}
	}
	return Formation{"Unknown", 0, 0, "none"}, false
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
	Formation{"Line", 0, 0, "all"},
	Formation{"AttColumn", 1, -1, "halfup"},
	Formation{"MarchColumn", 3, -2, "one"},
	Formation{"ClosedColumn", 5, -2, "half"},
	Formation{"Square", 9, -4, "one"},
	Formation{"Skirmish", -10, 0, "none"},
	Formation{"OpenOrder", -6, 0, "one"},
})
