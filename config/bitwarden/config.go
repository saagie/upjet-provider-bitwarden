package bitwarden

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure  database resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bitwarden_item_login", func(r *config.Resource) {
		r.UseAsync = true

		r.ShortGroup = "item.bitwarden"
	})
	p.AddResourceConfigurator("bitwarden_item_secure_note", func(r *config.Resource) {
		r.UseAsync = true

		r.ShortGroup = "item.bitwarden"
	})
}
