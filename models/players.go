package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Players struct {
	uadmin.Model
	Name  string `uadmin:"required;search;display_name:Player Name"`
	Photo string `uadmin:"image"`
	Team Team
	TeamID uint
	NameTeam string
	PPG  float64
	RPG float64
	APG float64
	PIE float64
	Birthdate time.Time
	BirthdateFormatted string `uadmin:"read_only"`
	Age int `uadmin:"read_only"`
	Experience int
	Height float64
	Weight int
	Country string
	Logo string `uadmin:"image"`
	Primary string
	Secondary string
}

func (p *Players) Save() {
	p.Age = ageCalculator(p.Birthdate)
	p.BirthdateFormatted = p.Birthdate.Format("01/02/06")
	uadmin.Preload(p)
	p.Logo = p.Team.Logo
	p.Primary = p.Team.Primary
	p.Secondary = p.Team.Secondary
	p.NameTeam = p.Team.Team
	uadmin.Save(p)
}

func ageCalculator (birthday time.Time) int {
	currentDate := time.Now()
	age := currentDate.Year() - birthday.Year()
	if currentDate.YearDay() <birthday.YearDay(){
		age--
	}
	return age
}