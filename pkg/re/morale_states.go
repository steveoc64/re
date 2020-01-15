package re

import "fyne.io/fyne/dataapi"

type MoraleState struct {
	Name             string
	SAFireModifier   int
	SATargetModifier int
}

func (m MoraleState) String() string {
	return m.Name
}

func (m MoraleState) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (m MoraleState) DeleteListener(int) {
}

func GetMoraleState(str string) (MoraleState, bool) {
	for _, v := range MoraleStates.Data {
		if m, ok := v.(MoraleState); ok {
			if m.Name == str {
				return m, true
			}
		}
	}
	return MoraleState{
		Name:             "Unknown",
		SAFireModifier:   0,
		SATargetModifier: 0,
	}, false
}

var MoraleStates = dataapi.NewSliceDataSource([]dataapi.DataItem{
	MoraleState{"Steady", 0, 0},
	MoraleState{"Steady", 0, 0},
	MoraleState{"Steady", 0, 0},
})
