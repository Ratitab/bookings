package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ratitab/bookings/pkg/config"
	"github.com/Ratitab/bookings/pkg/handlers"
	"github.com/Ratitab/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

// func Home(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "home.page.html")
// }

// func About(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "about.page.html")
// }

// func Devide(w http.ResponseWriter, r *http.Request) {
// }

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("err while parsing template", err)
// 		return
// 	}
// }

func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println(err)
	}

	app.TemplateCache = tc
	app.UseChace = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/devide", handlers.Repo.Devide)

	// _ = http.ListenAndServe(":8080", nil)
	serve := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
