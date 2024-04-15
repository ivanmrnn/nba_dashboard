package views

import (
    "net/http"
    "strings"

    "github.com/uadmin/uadmin"
)

// LoginHandler verifies login data and creates sessions for users.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Initialize the fields that we need in the custom struct.
    type Context struct {
        Err       string
        ErrExists bool
        Username  string
        Password  string
    }
    // Call the Context struct.
    c := Context{}

    // If the request method is POST
    if r.Method == "POST" {
        // This is a login request from the user.
        username := r.PostFormValue("username")
        username = strings.TrimSpace(strings.ToLower(username))
        password := r.PostFormValue("password")

        // Login using username and password.
        session, _ := uadmin.Login(r, username, password)

        // Check whether the session returned is nil or the user is not active.
        if session == nil || !session.User.Active {
            // Assign the login validation here that will be used for UI displaying.
            c.ErrExists = true
            c.Err = "Invalid username/password or inactive user"
        } else {
            // Create a session cookie
            cookie, _ := r.Cookie("session")
            if cookie == nil {
                cookie = &http.Cookie{}
            }
            cookie.Name = "session"
            cookie.Value = session.Key
            cookie.Path = "/"
            cookie.SameSite = http.SameSiteStrictMode
            http.SetCookie(w, cookie)

            // If valid, proceed to dashboard
            http.Redirect(w, r, "/nba_dashboard/", http.StatusSeeOther)
            return // Exit the function after redirection
        }
    }

    // Render the login filepath and pass the context data object to the HTML file.
    uadmin.RenderHTML(w, r, "templates/login.html", c)
}