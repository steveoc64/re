package re

import "fyne.io/fyne/dataapi"

type ContactSituation struct {
	Range      *dataapi.Int
	ReturnFire *dataapi.Bool
	Enfilade   *dataapi.Bool
	Weather    *dataapi.String
	Units      []*Unit
}

func NewSmallArmsSituation(units []*Unit) *ContactSituation {
	s := &ContactSituation{
		Range:      dataapi.NewInt(0),
		ReturnFire: dataapi.NewBool(true),
		Enfilade:   dataapi.NewBool(false),
		Weather:    dataapi.NewString("Clear"),
		Units:      units,
	}
	for _, v := range units {
		v.Situation = s
		v.CalcFF(0)
	}
	return s
}

func (s *ContactSituation) GetTarget(unitA *Unit) *Unit {
	for _, v := range s.Units {
		if v != unitA {
			return v
		}
	}
	return nil
}

func (s *ContactSituation) Clear() {
	for _, v := range s.Units {
		v.Clear()
	}
}

func (s *ContactSituation) Changed(string) {
	for _, v := range s.Units {
		v.Changed("")
	}
}
