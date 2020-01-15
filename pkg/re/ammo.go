package re

import "fyne.io/fyne/dataapi"

type AmmoState struct {
	Name             string
	SATargetModifier int
	SAFireModifier   int
}

func GetAmmoState(name string) (AmmoState, bool) {
	for _, v := range AmmoStates.Data {
		if f, ok := v.(AmmoState); ok {
			if f.Name == name {
				return f, true
			}
		}
	}
	return AmmoState{"Unknown", 0, 0}, false
}

func (a AmmoState) String() string {
	return a.Name
}

func (a AmmoState) AddListener(func(i dataapi.DataItem)) int {
	return 0
}

func (a AmmoState) DeleteListener(int) {
}

var AmmoStates = dataapi.NewSliceDataSource([]dataapi.DataItem{
	AmmoState{"FirstFire", 0, 5},
	AmmoState{"Good", 0, 0},
	AmmoState{"Depleted", 0, -6},
	AmmoState{"Exhausted", 0, -10},
})