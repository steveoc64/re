package re

import (
	"fmt"

	"fyne.io/fyne/dataapi"
)

type SmallArmsUnit struct {
	CloseOrderBases *dataapi.Int
	FiringBases     *dataapi.Int
	SupportingBases *dataapi.Int
	BasesDesc       *dataapi.String
	FireFactor      *dataapi.Int
	Hits            *dataapi.Int
	Class           *dataapi.String
	Range           *dataapi.Int
	Rifled          *dataapi.Bool
	Formation       *dataapi.String
	ClassList       *dataapi.SliceDataSource
}

func (s *SmallArmsUnit) CalcFF(f float64) {
	if s.FiringBases.Value() > s.CloseOrderBases.Value() {
		s.FiringBases.SetInt(s.CloseOrderBases.Value(), 0)

	}
	desc := fmt.Sprintf("%d/%d Bases firing", s.FiringBases.Value(), s.CloseOrderBases.Value())
	if s.SupportingBases.Value() > 0 {
		desc = fmt.Sprintf("%s, plus %d supports", desc, s.SupportingBases.Value())
	}
	s.BasesDesc.Set(desc, 0)
	s.FireFactor.SetInt(s.FiringBases.Value()*3+s.SupportingBases.Value(), 0)
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
		Range:           dataapi.NewInt(2),
		Rifled:          dataapi.NewBool(false),
		Formation:       dataapi.NewString("Line"),
		ClassList:       UnitClassesData,
	}
	s.CalcFF(0.0)
	return s
}
