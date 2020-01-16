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
	Die1D10         int
	Die2D10         int
	DieD6           int
	DieTotal        int
	DieMods         int
	AmmoState       *dataapi.String
	FireHitsAuto    int
	FireHitsExtra   int
	FireResults     *dataapi.String
	Terrain         *dataapi.String
	MoraleState     *dataapi.String
}

func (s *Unit) Changed(str string) {
	// TODO - dont change firing bases below every time - only do it when the formation changes
	println("something changed", str)
	s.FiringBases.SetInt(s.CloseOrderBases.Value())

	s.CalcFF(0)
	s.Situation.GetTarget(s).CalcFF(0)
}

func (s *Unit) CalcFF(f float64) {
	if s == nil {
		return
	}
	maxFiring := s.CloseOrderBases.Value()
	if f, ok := GetFormation(s.Formation); ok {
		switch f.MaxFire {
		//case "all":
		case "one":
			maxFiring = 1
		case "half":
			maxFiring = maxFiring / 2
		case "halfup":
			maxFiring = (maxFiring + 1) / 2
		case "none":
			maxFiring = 0
		}
	}
	if s.FiringBases.Value() > maxFiring {
		s.FiringBases.SetInt(maxFiring)
	}
	desc := fmt.Sprintf("%d of %d Bases firing", s.FiringBases.Value(), s.CloseOrderBases.Value())
	if s.SupportingBases.Value() > 0 {
		desc = fmt.Sprintf("%s, plus %d supports", desc, s.SupportingBases.Value())
	}
	s.BasesDesc.SetString(desc)

	rangeFactor := 1.0
	r := s.Situation.Range.Value()
	if r >= 12 { // 4-6 inches
		rangeFactor = 2.0 + (float64(r-12) / 3.0)
	} else if r >= 6 { // 2-4 inches
		rangeFactor = 1.0 + (float64(r-6) / 6.0)
	} // else go at 100%

	ff := s.FiringBases.Value()*3 + s.SupportingBases.Value()
	ff = int((float64(ff) / rangeFactor))
	s.FireFactor.SetInt(ff)

	s.calcDieMods()
}

func (s *Unit) calcDieMods() {
	dm := 0
	// start with the base unit class
	if c, ok := GetClassStats(s.Class); ok {
		dm += c.SAModifier
	}

	// target specific extra effects
	if tgt := s.Situation.GetTarget(s); tgt != nil {
		// formation
		if f, ok := GetFormation(tgt.Formation); ok {
			dm += f.SATargetModifier
		}
		// terrain
		if t, ok := GetTerrain(tgt.Terrain); ok {
			dm += t.SATargetModifier
		}
	}

	// ammo state effects
	if a, ok := GetAmmoState(s.AmmoState); ok {
		dm += a.SAFireModifier
	}

	// terrain effects
	if t, ok := GetTerrain(s.Terrain); ok {
		dm += t.SAFireModifier
	}

	// weather effects
	if w, ok := GetWeather(s.Situation.Weather); ok {
		dm += w.SAFireModifier
	}

	// morale state effects
	if m, ok := GetMoraleState(s.MoraleState); ok {
		dm += m.SAFireModifier
	}

	// current hits detract from FF
	dm += s.Hits.Value() * -1

	s.DieMods = dm
	s.DieModDesc.SetString(fmt.Sprintf("%+d Die Mod", dm))
}

// Clear the rolled state
func (s *Unit) Clear() {
	s.Die1D10 = 0
	s.Die2D10 = 0
	s.DieD6 = 0
	s.DieTotal = 0
	s.DieModDesc.SetString(fmt.Sprintf("%+d Die Mod", s.DieMods))
}

func (s *Unit) Roll() {
	d1 := rand.Intn(10) + 1
	d2 := rand.Intn(10) + 1
	d3 := rand.Intn(6) + 1
	dm := s.DieMods
	dt := d1 + d2 + dm
	s.Die1D10 = d1
	s.Die2D10 = d2
	s.DieD6 = d3
	s.DieTotal = dt
	s.DieModDesc.SetString(fmt.Sprintf("Rolled %d+%d with %+d = %d", d1, d2, dm, dt))
	ff := s.FireFactor.Value()

	s.FireHitsAuto, s.FireHitsExtra = calcSAFire(dt, ff)
	s.FireResults.SetString(fmt.Sprintf("(%d auto hits) + D6 = %d+ for extra hit", s.FireHitsAuto, s.FireHitsExtra))

	// Ammo depletion ?
	switch s.AmmoState.String() {
	case "FirstFire":
		s.AmmoState.SetString("Good")
		s.Changed("")
	case "Good":
		if d1 == 1 {
			s.AmmoState.SetString("Depleted")
			s.Changed("")
		}
	case "Depleted":
		if d1 <= 3 {
			s.AmmoState.SetString("Exhausted")
			s.Changed("")
		}
	}
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
		AmmoState:       dataapi.NewString("Good"),
		FireResults:     dataapi.NewString(""),
		Terrain:         dataapi.NewString("Open"),
	}
	return s
}
