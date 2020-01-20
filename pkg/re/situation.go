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
	victorTitle := ""

	if attacker.FireHitsTotal > defender.FireHitsTotal {
		winner = attacker
		loser = defender
		statusTitle = "Defender"
		victorTitle = "Attacker"
	} else if attacker.FireHitsTotal < defender.FireHitsTotal {
		winner = defender
		loser = attacker
		statusTitle = "Attacker"
		victorTitle = "Defender"
	} else {
		// inconclusive
		s.Status.SetString(fmt.Sprintf("After %d minutes of firefight", s.FirefightRound*20))
		return
	}

	mods = (loser.FireHitsTotal - winner.FireHitsTotal) * 2
	if loser.MoraleState.String() == "Shaken" {
		mods += 2
	}
	if loser.Hits.Value() >= loser.CloseOrderBases.Value() {
		mods += 2
	}
	switch loser.AmmoState.String() {
	case "Depleted", "Exhausted":
		mods += 5
	}

	d1 := rand.Intn(10) + 1
	d2 := rand.Intn(10) + 1
	println("die roll", d1, d2, mods, d1+d2+mods)
	v := d1 + d2 + mods
	switch {
	case v >= 23:
		loser.MoraleState.SetString("Broken")
		if winner.MoraleState.String() == "Eager" {
			s.Status.SetString(fmt.Sprintf("%s Breaks 8\" %s, %s Pursues with Bayonet", statusTitle, loser.MoraleState.String(), victorTitle))
			winner.MoraleCheckResult.SetString("Eager Pursuit with Bayonet")
			s.Resolved = true
		} else {
			s.SetStatus(loser)
		}
	case v >= 20:
		if loser.MoraleState.String() == "Steady" {
			loser.MoraleState.SetString("Disordered")
		}
		loser.MoraleCheckResult.SetString("Falls Back 5\"")
		if loser.MoraleState.String() == "Eager" {
			loser.MoraleCheckResult.SetString("Desperate Charge with Bayonet !")
			s.Status.SetString(fmt.Sprintf("%s Desperately Charges with Bayonet !", statusTitle))
		} else {
			if winner.MoraleState.String() == "Eager" && loser.MoraleState.String() != "Eager" {
				s.Status.SetString(fmt.Sprintf("%s Falls Back 5\" %s, %s Pursues with Bayonet", statusTitle, loser.MoraleState.String(), victorTitle))
				winner.MoraleCheckResult.SetString("Eager Pursuit with Bayonet")
			} else {
				s.Status.SetString(fmt.Sprintf("%s Falls Back 5\" %s", statusTitle, loser.MoraleState.String()))
			}
		}
		s.Resolved = true
	case v >= 17:
		if loser.MoraleState.String() == "Eager" {
			loser.MoraleCheckResult.SetString("Desperate Charge with Bayonet !")
			s.Status.SetString(fmt.Sprintf("%s Desperately Charges Enemy with Bayonet !", statusTitle))
		} else {
			loser.MoraleCheckResult.SetString("Falls Back 5\"")
			if winner.MoraleState.String() == "Eager" {
				winner.MoraleCheckResult.SetString("Eager Charge with Bayonet !")
				s.Status.SetString(fmt.Sprintf("%s Falls Back 5\" %s, %s Pursues with Bayonet !", statusTitle, loser.MoraleState.String(), victorTitle))
			} else {
				s.Status.SetString(fmt.Sprintf("%s Falls Back 5\" %s", statusTitle, loser.MoraleState.String()))
			}
			s.Resolved = true
		}
	case v >= 13:
		if loser.Terrain.String() == "Open" {
			if loser.MoraleState.String() == "Eager" {
				loser.MoraleCheckResult.SetString("Desperate Charge with Bayonet !")
				s.Status.SetString(fmt.Sprintf("%s Desperately Charges Enemy with Bayonet !", statusTitle))
			} else {
				loser.MoraleCheckResult.SetString("Falls Back 2\"")
				s.Status.SetString(fmt.Sprintf("%s Falls Back 2\" %s", statusTitle, loser.MoraleState.String()))
			}
		}
	default:
		s.Status.SetString(fmt.Sprintf("After %d minutes of firefight", s.FirefightRound*20))
	}
}
