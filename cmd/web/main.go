package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/zer0bi9/bookings/internal/config"
	"github.com/zer0bi9/bookings/internal/handlers"
	"github.com/zer0bi9/bookings/internal/render"
)

const portNumber = ":1323"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// var app config.AppConfig

	// change this to true when in production
	app.InProduction = false

	// Session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// Session - Store
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", repo.Home)
	// http.HandleFunc("/about", repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
