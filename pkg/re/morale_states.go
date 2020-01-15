package re

import (
	"fmt"
	"fyne.io/fyne/dataapi"
)

type MoraleState struct {
	Name           string
	SAFireModifier int
}

func (m MoraleState) String() string {
	return m.Name
}

func (m MoraleState) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (m MoraleState) DeleteListener(int) {
}

func GetMoraleState(str fmt.Stringer) (MoraleState, bool) {
	for _, v := range MoraleStates.Data {
		if m, ok := v.(MoraleState); ok {
			if m.Name == str.String() {
				return m, true
			}
		}
	}
	return MoraleState{
		Name:           "Unknown",
		SAFireModifier: 0,
	}, false
}

var MoraleStates = dataapi.NewSliceDataSource([]dataapi.DataItem{
	MoraleState{"Eager", 0},
	MoraleState{"Steady", 0},
	MoraleState{"Disordered", -5},
	MoraleState{"Shaken", -10},
	MoraleState{"Broken", -15},
})
