package re

import (
	"fmt"
	"fyne.io/fyne/dataapi"
)

type Weather struct {
	Name           string
	SAFireModifier int
}

func (w Weather) String() string {
	return w.Name
}

func (w Weather) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (w Weather) DeleteListener(int) {
}

func GetWeather(str fmt.Stringer) (Weather, bool) {
	for _, v := range Weathers.Data {
		if t, ok := v.(Weather); ok {
			if t.Name == str.String() {
				return t, true
			}
		}
	}
	return Weather{
		Name:           "Unknown",
		SAFireModifier: 0,
	}, false
}

var Weathers = dataapi.NewSliceDataSource([]dataapi.DataItem{
	Weather{Name: "Clear", SAFireModifier: 1},
	Weather{Name: "Damp", SAFireModifier: -1},
	Weather{Name: "LightRain", SAFireModifier: -2},
	Weather{Name: "HeavyRain", SAFireModifier: -5},
})
