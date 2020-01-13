package re

import "fyne.io/fyne/dataapi"

type Formation struct {
	Name        string
	SATargetMod int
}

func (f Formation) String() string {
	return f.Name
}

func (f Formation) AddListener(func(i dataapi.DataItem)) int {
	return 0
}

func (f Formation) DeleteListener(int) {
}

var Formations = []Formation{
	{"Line", 0},
	{"Mixed", 1},
	{"Closed Col", 5},
	{"Square", 9},
}

var FormationsData = dataapi.NewSliceDataSource()

func initFormatiotns() {
	println("init unit formations ")
	for _, v := range Formations {
		FormationsData.Append(v)
	}
}
