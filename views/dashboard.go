package views

import (
	"net/http"

	"github.com/ivanmrnn/nba_dashboard/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	type Context struct {
		User        string
		ProfilePhoto string `uadmin:"image"`

		Players []models.Players
		TotalPlayers int

		Team []models.Team
		TotalTeam int

		PlayerNames   []string 
		PlayerPhoto []string `uadmin:"image"`
		PlayerPPG  []float64
		PlayerRPG []float64
		PlayerAPG []float64
		PlayerPIE []float64
		PlayerBirthdateFormatted []string 
		PlayerAge []int 
		PlayerExperience []int
		PlayerHeight []float64
		PlayerWeight []int
		PlayerCountry []string
		PlayerLogo []string `uadmin:"image"`
		PlayerPrimary []string
		PlayerSecondary []string

		TeamNames   []string 
		TeamLogo []string `uadmin:"image"`
		TeamDivision   []string 
		TeamPrimary   []string 
		TeamSecondary   []string 
	}

	c := Context{}
	c.User = session.User.FirstName + " " + session.User.LastName
	c.ProfilePhoto = session.User.Photo 

	


	players := []models.Players{}
	uadmin.All(&players)
	c.Players = players

	c.TotalPlayers = uadmin.Count(&players, "id > 0")
	for _, player := range players {
        c.PlayerNames = append(c.PlayerNames, player.Name)
		c.PlayerPhoto = append(c.PlayerPhoto, player.Photo)
		c.PlayerPPG = append(c.PlayerPPG, player.PPG)
		c.PlayerRPG = append(c.PlayerRPG, player.RPG)
		c.PlayerAPG = append(c.PlayerAPG, player.APG)
		c.PlayerPIE = append(c.PlayerPIE, player.PIE)
		c.PlayerBirthdateFormatted = append(c.PlayerBirthdateFormatted, player.BirthdateFormatted)
		c.PlayerAge = append(c.PlayerAge, player.Age)
		c.PlayerExperience = append(c.PlayerExperience, player.Experience)
		c.PlayerHeight = append(c.PlayerHeight, player.Height)
		c.PlayerWeight = append(c.PlayerWeight, player.Weight)
		c.PlayerCountry = append(c.PlayerCountry, player.Country)
		c.PlayerLogo = append(c.PlayerLogo, player.Logo)
		c.PlayerPrimary = append(c.PlayerPrimary, player.Primary)
		c.PlayerSecondary = append(c.PlayerSecondary, player.Secondary)
    }


	team := []models.Team{}
	uadmin.All(&team)
	c.Team = team

	c.TotalTeam = uadmin.Count(&team, "id > 0")
	for _, team := range team {
		c.TeamNames = append(c.TeamNames, team.Name)
		c.TeamLogo = append(c.TeamLogo, team.Logo)
		c.TeamDivision = append(c.TeamDivision, team.Division)
		c.TeamPrimary = append(c.TeamPrimary, team.Primary)
		c.TeamSecondary = append(c.TeamSecondary, team.Secondary)
	}


	uadmin.RenderHTML(w, r, "templates/dashboard.html", c)
}
