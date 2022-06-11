// avoid include cycles by using this package to include external apps

package config

import "text/template"

// AppConfig hold sthe application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}

