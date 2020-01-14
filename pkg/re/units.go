package re

import (
	"math/rand"
	"time"

	"fyne.io/fyne/dataapi"
)

type ClassStats struct {
	ClassName  string
	SAModifier int
}

func (u ClassStats) String() string {
	return u.ClassName
}

func (u ClassStats) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (u ClassStats) DeleteListener(int) {
}

var UnitClasses = []ClassStats{
	{"Rabble", -6},
	{"Militia", -4},
	{"Landwehr", -3},
	{"Conscript", -2},
	{"Regular", -1},
	{"Veteran", 0},
	{"Crack Line", 2},
	{"Elite", 4},
	{"Grenadier", 6},
	{"Guard", 8},
	{"Old Guard", 10},
}

func GetClassStats(str string) (ClassStats,bool) {
	for _,v := range UnitClasses {
		if v.ClassName == str {
			return v,true
		}
	}
	return UnitClasses[0],false
}

var UnitClassesData = dataapi.NewSliceDataSource()

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	initUnits()
	initFormations()
}

func initUnits() {
	for _, v := range UnitClasses {
		UnitClassesData.Append(v)
	}
}
