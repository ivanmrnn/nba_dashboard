package models

import (
	"github.com/uadmin/uadmin"
)

type Team struct {
	uadmin.Model
	Name          string
	Team          string `uadmin:"read_only"`
	Logo          string `uadmin:"image"`
	Division      string
	Primary       string
	Secondary     string
}

func (p *Team) Save() {
	p.Team = p.Name
	uadmin.Save(p)
}
