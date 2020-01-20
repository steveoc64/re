package re

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/dataapi"
)

type Unit struct {
	Situation         *ContactSituation
	CloseOrderBases   *dataapi.Int
	FiringBases       *dataapi.Int
	SupportingBases   *dataapi.Int
	BasesDesc         *dataapi.String
	FireFactor        *dataapi.Int
	Hits              *dataapi.Int
	Class             *dataapi.String
	Rifled            *dataapi.Bool
	Formation         *dataapi.String
	DieModDesc        *dataapi.String
	Enfilade          *dataapi.Bool
	Die1D10           int
	Die2D10           int
	DieD6             int
	DieTotal          int
	DieMods           int
	AmmoState         *dataapi.String
	FireHitsAuto      int
	FireHitsExtra     int
	FireHitsTotal     int
	FireResults       *dataapi.String
	FireRolls         *dataapi.String
	FireTotal         *dataapi.String
	Terrain           *dataapi.String
	MoraleState       *dataapi.String
	MoraleCheckResult *dataapi.String
}

func (s *Unit) Changed(str string) {
	s.CalcFF(0)
	s.Situation.GetTarget(s).CalcFF(0)
}

func (s *Unit) HitsChanged(str string) {
	s.CalcFF(0)
	s.Situation.GetTarget(s).CalcFF(0)
}

func (s *Unit) FormationChanged(str string) {
	s.FiringBases.SetInt(s.CloseOrderBases.Value())
	s.Changed(str)
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
	} else if r == 0 {
		rangeFactor = 0.6
	} // else go at 100%

	// Add supporting bases
	ff := s.FiringBases.Value()*3 + s.SupportingBases.Value()
	ff = int((float64(ff) / rangeFactor))

	// Take hits
	ff = ff - s.Hits.Value()

	// enfilade ?
	if s.Enfilade.Value() {
		ff = ff + 5
	}

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

	// terrain effects firing out from
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

	// current hits detract from FF as well as negatively affect die roll
	dm += s.Hits.Value() * -1

	// completely fresh unit with no hits at all
	if s.Hits.Value() == 0 && s.MoraleState.String() == "Steady" {
		dm += 3
	}

	s.DieMods = dm
	s.DieModDesc.SetString(fmt.Sprintf("%+d", dm))
}

// Clear the rolled state
func (s *Unit) Clear() {
	s.Die1D10 = 0
	s.Die2D10 = 0
	s.DieD6 = 0
	s.DieTotal = 0
	s.DieModDesc.SetString(fmt.Sprintf("%+d", s.DieMods))
	s.FireRolls.SetString("")
	s.FireTotal.SetString("")
	s.FireResults.SetString("")
	s.MoraleState.SetString("Steady")
	s.Hits.SetInt(0)
	s.CloseOrderBases.SetInt(4)
	s.FiringBases.SetInt(4)
	s.SupportingBases.SetInt(0)
	s.Formation.SetString("Line")
	s.AmmoState.SetString("Good")
	s.MoraleCheckResult.SetString("")
}

func (s *Unit) Fire() {
	d1 := rand.Intn(10) + 1
	d2 := rand.Intn(10) + 1
	d3 := rand.Intn(6) + 1
	dm := s.DieMods
	dt := d1 + d2 + dm
	s.Die1D10 = d1
	s.Die2D10 = d2
	s.DieD6 = d3
	s.DieTotal = dt
	roll := fmt.Sprintf("2D10[%d+%d] %+d = %d", d1, d2, dm, dt)

	ammoEffect := ""

	ff := s.FireFactor.Value()

	s.FireHitsAuto, s.FireHitsExtra = calcSAFire(dt, ff)

	if s.Situation.AutoDice.Value() {
		d6 := rand.Intn(6) + 1
		s.FireHitsTotal = s.FireHitsAuto
		if s.FireHitsExtra > 0 && d6 >= s.FireHitsExtra {
			s.FireHitsTotal++
		}
		roll = roll + fmt.Sprintf(" / D6[%d]", d6)
		ftotal := fmt.Sprintf("%d Hit", s.FireHitsTotal)
		if s.FireHitsTotal != 1 {
			ftotal = ftotal + "s"
		}
		s.FireTotal.SetString(ftotal)

		// auto apply to opponent
		if target := s.Situation.GetTarget(s); target != nil {
			target.Hits.SetInt(target.Hits.Value() + s.FireHitsTotal)
			target.MoraleCheck()
		}
	} // else add a pair of buttons for the user to select from after rolling dice
	s.FireRolls.SetString(roll)

	// Ammo depletion ?
	switch s.AmmoState.String() {
	case "FirstFire":
		s.AmmoState.SetString("Good")
		s.Changed("")
	case "Good":
		if d1 == 1 {
			s.AmmoState.SetString("Depleted")
			s.Changed("")
			ammoEffect = " Depleted"
		}
	case "Depleted":
		if d1 <= 3 {
			s.AmmoState.SetString("Exhausted")
			s.Changed("")
			ammoEffect = " Exhausted"
		}
	}

	res := fmt.Sprintf("%d / %d+ Hits%s", s.FireHitsAuto, s.FireHitsExtra, ammoEffect)
	s.FireResults.SetString(res)
}

func (s *Unit) MoraleCheck() bool {
	// get the base number to fail
	if m, ok := GetClassStats(s.Class); ok {
		mods := 0

		// get current morale state
		if ms, ok := GetMoraleState(s.MoraleState); ok {
			mods += ms.MoraleMod
		}

		// is in line
		switch s.Formation.String() {
		case "Line":
			mods += m.MoraleLineMod
		case "Square":
			mods += 3
		case "ClosedColumn":
			mods += 1
		}

		// apply terrain effects
		if t, ok := GetTerrain(s.Terrain); ok {
			mods += t.MoraleMod
		}

		// get the losses
		lossPercent := (s.Hits.Value() * 100) / (s.CloseOrderBases.Value() * 3)
		if s.Hits.Value() >= (s.CloseOrderBases.Value() * 3) {
			// totally lost
			s.MoraleCheckResult.SetString("Unit Destroyed")
			s.Situation.SetStatus(s)
			return false
		}

		d6 := rand.Intn(6) + 1
		if s.Hits.Value() >= (s.CloseOrderBases.Value() * 3 / 2) {
			// over 50% losses, rounding up in favour of the unit
			if s.MoraleState.String() == "Steady" && d6 <= 3 {
				s.MoraleState.SetString("Shaken")
			}
			mods -= 7
		} else if s.Hits.Value() >= s.CloseOrderBases.Value() {
			// as many hits as there are bases = 1/3rd losses
			mods -= 4
			if s.MoraleState.String() == "Steady" && d6 <= 2 {
				s.MoraleState.SetString("Disordered")
			}
		} else {
			if m.ID >= 5 {
				s.MoraleCheckResult.SetString("Holds Steady")
				return true
			}
		}

		d1 := rand.Intn(10) + 1
		d2 := rand.Intn(10) + 1
		println("die rolls", d1, d2, "=", d1+d2, "modded to", d1+d2+mods, "fails on", m.MoraleCheck)
		if d1+d2+mods <= m.MoraleCheck {
			s.MoraleState.SetString("Broken")
			s.MoraleCheckResult.SetString(fmt.Sprintf("Breaks with %d%% Losses", lossPercent))
			s.Situation.SetStatus(s)
			return false
		}

		s.MoraleCheckResult.SetString(fmt.Sprintf("Holds with %d%% Losses", lossPercent))
		return true
	}
	return false
}

func NewSmallArmsUnit() *Unit {
	s := &Unit{
		CloseOrderBases:   dataapi.NewInt(4),
		FiringBases:       dataapi.NewInt(4),
		SupportingBases:   dataapi.NewInt(0),
		BasesDesc:         dataapi.NewString(""),
		FireFactor:        dataapi.NewInt(12),
		Hits:              dataapi.NewInt(0),
		Class:             dataapi.NewString("Regular"),
		MoraleState:       dataapi.NewString("Steady"),
		Rifled:            dataapi.NewBool(false),
		Formation:         dataapi.NewString("Line"),
		DieModDesc:        dataapi.NewString(""),
		AmmoState:         dataapi.NewString("Good"),
		FireResults:       dataapi.NewString(""),
		FireRolls:         dataapi.NewString(""),
		FireTotal:         dataapi.NewString(""),
		Terrain:           dataapi.NewString("Open"),
		Enfilade:          dataapi.NewBool(false),
		MoraleCheckResult: dataapi.NewString(""),
	}
	s.Hits.AddListener(func(dataapi.DataItem) { s.CalcFF(0) })
	s.CloseOrderBases.AddListener(func(item dataapi.DataItem) { s.CalcFF(0) })
	s.FiringBases.AddListener(func(item dataapi.DataItem) { s.CalcFF(0) })
	s.SupportingBases.AddListener(func(item dataapi.DataItem) { s.CalcFF(0) })
	s.Rifled.AddListener(func(item dataapi.DataItem) { s.CalcFF(0) })
	s.Enfilade.AddListener(func(item dataapi.DataItem) { s.CalcFF(0) })
	return s
}
