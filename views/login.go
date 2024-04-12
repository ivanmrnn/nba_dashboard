package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// LoginHandler verifies login data and creating sessions for users.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the fields that we need in the custom struct.
	type Context struct {
		Err         string
		ErrExists   bool
		OTPRequired bool
		Username    string
		Password    string
	}
	// Call the Context struct.
	c := Context{}

	// If the request method is POST
	if r.Method == "POST" {
		// This is a login request from the user.
		username := r.PostFormValue("username")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.PostFormValue("password")
		otp := r.PostFormValue("otp")

		// Login2FA login using username, password and otp for users with OTPRequired = true.
		session := uadmin.Login2FA(r, username, password, otp)

		// Check whether the session returned is nil or the user is not active.
		if session == nil || !session.User.Active {
			/* Assign the login validation here that will be used for UI displaying. ErrExists and
			   Err fields are coming from the Context struct. */
			c.ErrExists = true
			c.Err = "Invalid username/password or inactive user"

		} else {
			
			cookie, _ := r.Cookie("session")
			if cookie == nil {
				cookie = &http.Cookie{}
			}
			cookie.Name = "session"
			cookie.Value = session.Key
			cookie.Path = "/"
			cookie.SameSite = http.SameSiteStrictMode
			http.SetCookie(w, cookie)

			//if valid proceed to dashboard
			http.Redirect(w, r, "/nba_dashboard/", http.StatusSeeOther)
		}
	}

	// Render the login filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/login.html", c)
}
