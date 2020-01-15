package re

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/dataapi"
)

type Unit struct {
	Situation       *ContactSituation
	CloseOrderBases *dataapi.Int
	FiringBases     *dataapi.Int
	SupportingBases *dataapi.Int
	BasesDesc       *dataapi.String
	FireFactor      *dataapi.Int
	Hits            *dataapi.Int
	Class           *dataapi.String
	Rifled          *dataapi.Bool
	Formation       *dataapi.String
	DieModDesc      *dataapi.String
	Die1D10         *dataapi.Int
	Die2D10         *dataapi.Int
	DieD6           *dataapi.Int
	DieTotal        *dataapi.Int
	DieMods         *dataapi.Int
	AmmoState       *dataapi.String
}

// ClassChanged handler
func (s *Unit) ClassChanged(str string) {
	s.CalcFF(0.0)
}

func (s *Unit) Changed(str string) {
	s.CalcFF(0)
	s.Situation.GetTarget(s).CalcFF(0)
}

func (s *Unit) CalcFF(f float64) {
	if s == nil {
		return
	}
	if s.FiringBases.Value() > s.CloseOrderBases.Value() {
		s.FiringBases.SetInt(s.CloseOrderBases.Value(), 0)

	}
	desc := fmt.Sprintf("%d of %d Bases firing", s.FiringBases.Value(), s.CloseOrderBases.Value())
	if s.SupportingBases.Value() > 0 {
		desc = fmt.Sprintf("%s, plus %d supports", desc, s.SupportingBases.Value())
	}
	s.BasesDesc.Set(desc, 0)

	rangeFactor := 1.0
	r := s.Situation.Range.Value()
	if r >= 12 { // 4-6 inches
		rangeFactor = 2.0 + (float64(r-12) / 3.0)
	} else if r >= 6 { // 2-4 inches
		rangeFactor = 1.0 + (float64(r-6) / 6.0)
	} // else go at 100%

	ff := s.FiringBases.Value()*3 + s.SupportingBases.Value()
	ff = int((float64(ff) / rangeFactor))
	s.FireFactor.SetInt(ff, 0)

	s.calcDieMods()
}

func (s *Unit) calcDieMods() {
	dm := 0
	if c, ok := GetClassStats(s.Class.String()); ok {
		dm += c.SAModifier
	}

	if t := s.Situation.GetTarget(s); t != nil {
		if f, ok := GetFormation(t.Formation.String()); ok {
			dm += f.SATargetModifier
		}
	}

	if a, ok := GetAmmoState(s.AmmoState.String()); ok {
		dm += a.SAFireModifier
	}
	s.DieMods.SetInt(dm, 0)
	s.DieModDesc.Set(fmt.Sprintf("%+d Die Mod", dm), 0)
}

// Clear the rolled state
func (s *Unit) Clear() {
	s.Die1D10.SetInt(0, 0)
	s.Die2D10.SetInt(0, 0)
	s.DieD6.SetInt(0, 0)
	s.DieTotal.SetInt(0, 0)
	s.DieModDesc.Set(fmt.Sprintf("%+d Die Mod", s.DieMods.Value()), 0)
}

func (s *Unit) Roll() {
	d1 := rand.Intn(10) + 1
	d2 := rand.Intn(10) + 1
	d3 := rand.Intn(6) + 1
	dm := s.DieMods.Value()
	dt := d1 + d2 + dm
	s.Die1D10.SetInt(d1, 0)
	s.Die2D10.SetInt(d2, 0)
	s.DieD6.SetInt(d3, 0)
	s.DieTotal.SetInt(dt, 0)
	s.DieModDesc.Set(fmt.Sprintf("Rolled %d+%d / %d with mods %+d = %d", d1, d2, d3, dm, dt), 0)
	ff := s.FireFactor.Value()

	band := int(dt / 4)
	if band < 0 {
		band = 0
	}
	println("die band is", band)
	basePips := 0
	pip := 1.0
	switch {
	case ff > 20:
		pip = 4.5
	case ff > 16:
		pip = 3.8
	case ff > 14:
		pip = 3.0
	case ff > 12:
		pip = 2.6
	case ff > 10:
		pip = 2.4
	case ff > 8:
		pip = 2.2
	case ff > 6:
		pip = 2.0
	case ff > 4:
		pip = 1.8
	case ff > 3:
		pip = 1.6
	case ff > 2:
		pip = 1.4
	default:
		pip = 1.0
	}
	totalPips := basePips + int(pip*float64(band))
	println("band", band, "pippage", pip, "pips", totalPips)
}

func NewSmallArmsUnit() *Unit {
	s := &Unit{
		CloseOrderBases: dataapi.NewInt(4),
		FiringBases:     dataapi.NewInt(4),
		SupportingBases: dataapi.NewInt(0),
		BasesDesc:       dataapi.NewString(""),
		FireFactor:      dataapi.NewInt(12),
		Hits:            dataapi.NewInt(0),
		Class:           dataapi.NewString("Regular"),
		Rifled:          dataapi.NewBool(false),
		Formation:       dataapi.NewString("Line"),
		DieModDesc:      dataapi.NewString(""),
		Die1D10:         dataapi.NewInt(0),
		Die2D10:         dataapi.NewInt(0),
		DieD6:           dataapi.NewInt(0),
		DieTotal:        dataapi.NewInt(0),
		DieMods:         dataapi.NewInt(0),
		AmmoState:       dataapi.NewString("Good"),
	}
	return s
}
