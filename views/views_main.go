package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	path := strings.TrimSuffix(r.URL.Path, "/")

	//check if user is authenticated, if not session=nil
	session := uadmin.IsAuthenticated(r)

	//url path
	switch path {
	case "login":
		LoginHandler(w, r)
	case "nba_dashboard":
		//if /nba_dashboard is entered and no session, to be directed to login
		if session == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			DashboardHandler(w, r, session)
		}
	case "logout":
		LogoutHandler(w, r, session)
	default:
		w.Write([]byte("Not Found"))
	}
}
