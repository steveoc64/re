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

func GetClassStats(str string) (ClassStats, bool) {
	for _, v := range UnitClasses.Data {
		if c, ok := v.(ClassStats); ok {
			if c.ClassName == str {
				return c, true
			}
		}
	}
	return ClassStats{
		ClassName:  "Unknown",
		SAModifier: 0,
	}, false
}

var UnitClasses = dataapi.NewSliceDataSource([]dataapi.DataItem{
	ClassStats{"Rabble", -6},
	ClassStats{"Militia", -4},
	ClassStats{"Landwehr", -3},
	ClassStats{"Conscript", -2},
	ClassStats{"Regular", -1},
	ClassStats{"Veteran", 0},
	ClassStats{"Crack Line", 2},
	ClassStats{"Elite", 4},
	ClassStats{"Grenadier", 6},
	ClassStats{"Guard", 8},
	ClassStats{"Old Guard", 10},
})

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
