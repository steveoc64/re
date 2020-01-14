package re

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/dataapi"
)

type SmallArmsSituation struct {
	Range      *dataapi.Int
	ReturnFire *dataapi.Bool
	Units      []*SmallArmsUnit
}

func NewSmallArmsSituation(units []*SmallArmsUnit) *SmallArmsSituation {
	s := &SmallArmsSituation{
		Range:      dataapi.NewInt(0),
		ReturnFire: dataapi.NewBool(true),
		Units:      units,
	}
	for _, v := range units {
		v.Situation = s
		v.CalcFF(0)
	}
	return s
}

func (s *SmallArmsSituation) GetTarget(unitA *SmallArmsUnit) *SmallArmsUnit {
	for _,v := range s.Units {
		if v != unitA {
			return v
		}
	}
	return nil
}

func (s *SmallArmsSituation) Roll() {
	for _,v := range s.Units {
		v.Roll()
	}
}

type SmallArmsUnit struct {
	Situation       *SmallArmsSituation
	CloseOrderBases *dataapi.Int
	FiringBases     *dataapi.Int
	SupportingBases *dataapi.Int
	BasesDesc       *dataapi.String
	FireFactor      *dataapi.Int
	Hits            *dataapi.Int
	Class           *dataapi.String
	ClassList       *dataapi.SliceDataSource
	Rifled          *dataapi.Bool
	Formation       *dataapi.String
	FormationList   *dataapi.SliceDataSource
	DieModDesc      *dataapi.String
	Die1D10         *dataapi.Int
	Die2D10         *dataapi.Int
	DieD6           *dataapi.Int
	DieTotal        *dataapi.Int
	DieMods         *dataapi.Int
}

// ClassChanged handler
func (s *SmallArmsUnit) ClassChanged(str string) {
	s.CalcFF(0.0)
}

func (s *SmallArmsUnit) FormationChanged(str string) {
	s.Situation.GetTarget(s).CalcFF(0)
}

func (s *SmallArmsUnit) CalcFF(f float64) {
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

func (s *SmallArmsUnit) calcDieMods() {
	dm := 0
	if c, ok := GetClassStats(s.Class.String()); ok {
		dm += c.SAModifier
	}

	if t := s.Situation.GetTarget(s); t != nil {
		if f,ok := GetFormation(t.Formation.String()); ok {
			dm += f.SATargetMod
		}
	}
	s.DieMods.SetInt(dm, 0)
	s.DieModDesc.Set(fmt.Sprintf("%+d Die Mod", dm), 0)
}

func (s *SmallArmsUnit) Roll() {
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
		pip = 1.6
	case ff > 16:
		pip = 1.5
	case ff > 14:
		pip =
	case ff > 12:
		pip = 2.6
	case ff > 10:
		pip = 2.4
	case ff > 8:
		pip = 2.2
	case ff > 6:
		pip = 2.0
	case ff > 4:
		pip = 0.9
	case ff > 3:
		pip = 0.8
	case ff > 2:
		pip = 0.7
	default:
		pip = 0.6
	}
	totalPips := basePips + int(pip * float64(band))
	println("band",band,"pippage", pip, "pips", totalPips)
}

func NewSmallArmsUnit() *SmallArmsUnit {
	s := &SmallArmsUnit{
		CloseOrderBases: dataapi.NewInt(4),
		FiringBases:     dataapi.NewInt(4),
		SupportingBases: dataapi.NewInt(0),
		BasesDesc:       dataapi.NewString(""),
		FireFactor:      dataapi.NewInt(12),
		Hits:            dataapi.NewInt(0),
		Class:           dataapi.NewString("Regular"),
		ClassList:       UnitClassesData,
		Rifled:          dataapi.NewBool(false),
		Formation:       dataapi.NewString("Line"),
		FormationList:   FormationsData,
		DieModDesc:      dataapi.NewString(""),
		Die1D10:         dataapi.NewInt(0),
		Die2D10:         dataapi.NewInt(0),
		DieD6:           dataapi.NewInt(0),
		DieTotal:        dataapi.NewInt(0),
		DieMods:         dataapi.NewInt(0),
	}
	return s
}
