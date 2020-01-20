package re

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/dataapi"
)

type ContactSituation struct {
	Range          *dataapi.Int
	ReturnFire     *dataapi.Bool
	Enfilade       *dataapi.Bool
	AutoDice       *dataapi.Bool
	Weather        *dataapi.String
	Status         *dataapi.String
	FirefightRound int
	Resolved       bool
	Units          []*Unit
}

func NewSmallArmsSituation(units []*Unit) *ContactSituation {
	s := &ContactSituation{
		Range:          dataapi.NewInt(2),
		ReturnFire:     dataapi.NewBool(true),
		Enfilade:       dataapi.NewBool(false),
		AutoDice:       dataapi.NewBool(true),
		Weather:        dataapi.NewString("Clear"),
		Status:         dataapi.NewString(""),
		Units:          units,
		FirefightRound: 0,
		Resolved:       false,
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
	s.Range.SetInt(2)
	s.ReturnFire.SetBool(true)
	s.Enfilade.SetBool(false)
	s.Status.SetString("")
	s.FirefightRound = 0
	s.Resolved = false
	for _, v := range s.Units {
		v.Clear()
	}
}

func (s *ContactSituation) Changed(string) {
	for _, v := range s.Units {
		v.Changed("")
	}
}

func (s *ContactSituation) SetStatus(u *Unit) {
	uu := "Attacker"
	if u != s.Units[0] {
		uu = "Defender"
	}
	s.Status.SetString(fmt.Sprintf("%s %s", uu, u.MoraleCheckResult.String()))
	s.Resolved = true
}

func (s *ContactSituation) FirefightCheck() {
	if s.Resolved {
		return
	}
	println("Firefight Check")

	attacker := s.Units[0]
	defender := s.Units[1]
	s.FirefightRound++
	mods := 0
	var winner, loser *Unit
	statusTitle := ""

	if attacker.FireHitsTotal > defender.FireHitsTotal {
		winner = attacker
		loser = defender
		statusTitle = "Defender"
	} else if attacker.FireHitsTotal < defender.FireHitsTotal {
		winner = defender
		loser = attacker
		statusTitle = "Attacker"
	} else {
		// inconclusive
		s.Status.SetString(fmt.Sprintf("After %d minutes of firefight", s.FirefightRound*20))
		return
	}

	mods = (loser.FireHitsTotal-winner.FireHitsTotal)*2
	if loser.MoraleState.String() == "Shaken" {
		mods += 2
	}
	if loser.Hits.Value() >= loser.CloseOrderBases.Value() {
		mods += 2
	}
	switch loser.AmmoState.String() {
	case "Depleted","Exhausted":
		mods += 5
	}

	d1 := rand.Intn(10)+1
	d2 := rand.Intn(10)+1
	println("die roll", d1, d2, mods, d1+d2+mods)
	v := d1+d2+mods
	switch {
	case v >= 23:
		loser.MoraleState.SetString("Broken")
		s.SetStatus(loser)
	case v >= 20:
		loser.MoraleState.SetString("Disordered")
		loser.MoraleCheckResult.SetString("Falls Back 5\" Disordered")
		s.Status.SetString(fmt.Sprintf("%s Falls Back 5\" Disordered", statusTitle))
		s.Resolved = true
	case v >= 17:
		loser.MoraleState.SetString("Steady")
		loser.MoraleCheckResult.SetString("Falls Back 5\"")
		s.Status.SetString(fmt.Sprintf("%s Falls Back 5\" in Good Order", statusTitle))
		s.Resolved = true
	case v >= 13:
		if loser.Terrain.String() == "Open" {
			loser.MoraleCheckResult.SetString("Falls Back 2\"")
			s.Status.SetString(fmt.Sprintf("%s Falls Back 2\" in Good Order", statusTitle))
		}
	default:
		s.Status.SetString(fmt.Sprintf("After %d minutes of firefight", s.FirefightRound*20))
	}
}
