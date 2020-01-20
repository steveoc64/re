package re

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/dataapi"
)

type ClassStats struct {
	ClassName  string
	ID int
	SAModifier int
	MoraleCheck int
	MoraleLineMod int
}

func (u ClassStats) String() string {
	return u.ClassName
}

func (u ClassStats) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (u ClassStats) DeleteListener(int) {
}

func GetClassStats(str fmt.Stringer) (ClassStats, bool) {
	for _, v := range UnitClasses.Data {
		if c, ok := v.(ClassStats); ok {
			if c.ClassName == str.String() {
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
	ClassStats{"Rabble",1, -6, 9, -3},
	ClassStats{"Militia",2, -4, 7, -3},
	ClassStats{"Landwehr",3, -3, 6, -3},
	ClassStats{"Conscript",4, -2, 5, -1},
	ClassStats{"Regular",5, -1, 4, -1},
	ClassStats{"Veteran",6, 0, 3, 0},
	ClassStats{"Crack Line",7, 2, 2, 0},
	ClassStats{"Elite",8, 4, 1, 0},
	ClassStats{"Grenadier",9, 6, 0, 0},
	ClassStats{"Guard",10, 8, -1, 0},
	ClassStats{"Old Guard",11, 10, -3, 0},
})

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
