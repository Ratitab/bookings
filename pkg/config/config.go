package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseChace      bool
	TemplateCache map[string]*template.Template
	// InfoLog       *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}
