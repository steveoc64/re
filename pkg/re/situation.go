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
	TotalLosses    int
	Attacker       *Unit
	Defender       *Unit
}

func (s *ContactSituation) Copy(sit *ContactSituation) {
	s.Range.SetInt(sit.Range.Value())
	s.ReturnFire.SetBool(sit.ReturnFire.Value())
	s.Enfilade.SetBool(sit.Enfilade.Value())
	s.AutoDice.SetBool(sit.AutoDice.Value())
	s.Weather.SetString(sit.Weather.String())
	s.Status.SetString(sit.Status.String())
	s.FirefightRound = sit.FirefightRound
	s.Resolved = sit.Resolved
	s.Attacker.Copy(sit.Attacker)
	s.Defender.Copy(sit.Defender)
	s.TotalLosses = sit.TotalLosses
}

func NewSmallArmsSituation() *ContactSituation {
	s := &ContactSituation{
		Range:          dataapi.NewInt(2),
		ReturnFire:     dataapi.NewBool(true),
		Enfilade:       dataapi.NewBool(false),
		AutoDice:       dataapi.NewBool(true),
		Weather:        dataapi.NewString("Clear"),
		Status:         dataapi.NewString(""),
		Attacker:       NewUnit(),
		Defender:       NewUnit(),
		FirefightRound: 0,
		Resolved:       false,
		TotalLosses:    0,
	}
	s.Attacker.Situation = s
	s.Defender.Situation = s
	s.Attacker.CalcFF(0)
	s.Defender.CalcFF(0)
	return s
}

func (s *ContactSituation) GetTarget(unitA *Unit) *Unit {
	if unitA == s.Attacker {
		return s.Defender
	}
	if unitA == s.Defender {
		return s.Attacker
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
	s.Attacker.Clear()
	s.Defender.Clear()
	s.TotalLosses = 0
}

func (s *ContactSituation) Changed(string) {
	s.Attacker.Changed("")
	s.Defender.Changed("")
}

func (s *ContactSituation) SetStatus(u *Unit) {
	uu := "Attacker"
	if u != s.Attacker {
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

	attacker := s.Attacker
	defender := s.Defender
	s.FirefightRound++
	mods := 0
	var winner, loser *Unit
	statusTitle := ""
	victorTitle := ""
	aHits := attacker.FireHitsTotal - attacker.SupportingBases.Value()
	dHits := defender.FireHitsTotal - defender.SupportingBases.Value()
	if aHits < 0 {
		s.TotalLosses = attacker.FireHitsTotal * 20
		aHits = 0
	}
	if dHits < 0 {
		s.TotalLosses += defender.FireHitsTotal * 20
		dHits = 0
	}
	s.TotalLosses = s.TotalLosses + 60*aHits + 60*dHits
	statusString := fmt.Sprintf("After %d minutes and %d men down", s.FirefightRound*20, s.TotalLosses)

	if aHits > dHits {
		winner = attacker
		loser = defender
		statusTitle = "Defender"
		victorTitle = "Attacker"
	} else if aHits < dHits {
		winner = defender
		loser = attacker
		statusTitle = "Attacker"
		victorTitle = "Defender"
	} else {
		// inconclusive
		s.Status.SetString(fmt.Sprintf("%s, firefight stalemate", statusString))
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
			s.Status.SetString(fmt.Sprintf("%s, %s Breaks 8\" %s, %s Pursues with Bayonet", statusString, statusTitle, loser.MoraleState.String(), victorTitle))
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
			s.Status.SetString(fmt.Sprintf("%s, %s Desperately Charges with Bayonet !", statusString, statusTitle))
		} else {
			if winner.MoraleState.String() == "Eager" && loser.MoraleState.String() != "Eager" {
				s.Status.SetString(fmt.Sprintf("%s Falls Back 5\" %s, %s Pursues with Bayonet", statusTitle, loser.MoraleState.String(), victorTitle))
				winner.MoraleCheckResult.SetString(fmt.Sprintf("%s, Eager Pursuit with Bayonet", statusString))
			} else {
				s.Status.SetString(fmt.Sprintf("%s, %s Falls Back 5\" %s", statusString, statusTitle, loser.MoraleState.String()))
			}
		}
		s.Resolved = true
	case v >= 17:
		if loser.MoraleState.String() == "Eager" {
			loser.MoraleCheckResult.SetString("Desperate Charge with Bayonet !")
			s.Status.SetString(fmt.Sprintf("%s, %s Desperately Charges Enemy with Bayonet !", statusString, statusTitle))
		} else {
			loser.MoraleCheckResult.SetString("Falls Back 5\"")
			if winner.MoraleState.String() == "Eager" {
				winner.MoraleCheckResult.SetString("Eager Charge with Bayonet !")
				s.Status.SetString(fmt.Sprintf("%s, %s Falls Back 5\" %s, %s Pursues with Bayonet !", statusString, statusTitle, loser.MoraleState.String(), victorTitle))
			} else {
				s.Status.SetString(fmt.Sprintf("%s, %s Falls Back 5\" %s", statusString, statusTitle, loser.MoraleState.String()))
			}
			s.Resolved = true
		}
	case v >= 13:
		if loser.Terrain.String() == "Open" {
			if loser.MoraleState.String() == "Eager" {
				loser.MoraleCheckResult.SetString("Desperate Charge with Bayonet !")
				s.Status.SetString(fmt.Sprintf("%s, %s Desperately Charges Enemy with Bayonet !", statusString, statusTitle))
			} else {
				loser.MoraleCheckResult.SetString("Falls Back 2\"")
				s.Status.SetString(fmt.Sprintf("%s, %s Falls Back 2\" %s", statusString, statusTitle, loser.MoraleState.String()))
			}
		}
	default:
		s.Status.SetString(fmt.Sprintf("%s, firefight continues", statusString))
	}
}
