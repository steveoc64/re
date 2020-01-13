package re

import "fyne.io/fyne/dataapi"

type UnitStats struct {
	ClassName string
	SAModifier int
}

func (u UnitStats) String() string {
	return u.ClassName
}

func (u UnitStats) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (u UnitStats) DeleteListener(int) {
}

var UnitClasses = []UnitStats{
	{"Rabble", -6},
	{"Militia", -4},
	{"Landwehr", -3},
	{"Conscript", -2},
	{"Regular", -1,},
	{"Veteran", 0},
	{"Crack Line", 2},
	{"Elite", 4},
	{"Grenadier", 6},
	{"Guard", 8},
	{"Old Guard", 10},
}

var UnitClassesData = dataapi.NewSliceDataSource()

func init() {
	println("init unit clasess ")
	for _,v := range UnitClasses {
		UnitClassesData.Append(v)
	}
}
