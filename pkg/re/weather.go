package re

import "fyne.io/fyne/dataapi"

type Weather struct {
	Name             string
	SAFireModifier   int
	SATargetModifier int
}

func (w Weather) String() string {
	return w.Name
}

func (w Weather) AddListener(func(dataapi.DataItem)) int {
	return 0
}

func (w Weather) DeleteListener(int) {
}

func GetWeather(str string) (Weather, bool) {
	for _, v := range Weathers.Data {
		if t, ok := v.(Weather); ok {
			if t.Name == str {
				return t, true
			}
		}
	}
	return Weather{
		Name:             "Unknown",
		SAFireModifier:   0,
		SATargetModifier: 0,
	}, false
}

var Weathers = dataapi.NewSliceDataSource([]dataapi.DataItem{
	Weather{Name: "Clear", SAFireModifier: 0, SATargetModifier: 0},
	Weather{Name: "LightRain", SAFireModifier: 0, SATargetModifier: 0},
	Weather{Name: "HeavyRain", SAFireModifier: 0, SATargetModifier: 0},
})
