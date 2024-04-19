package main

import (
	"net/http"

	"github.com/ivanmrnn/nba_dashboard/models"
	"github.com/ivanmrnn/nba_dashboard/views"
	"github.com/uadmin/uadmin"
)

func main() {
	//register models
	uadmin.Register(
		models.Players{},
		models.Team{},
	)
	
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "NBA Players"

	http.HandleFunc("/", uadmin.Handler(views.MainHandler))

	uadmin.StartServer()

}
