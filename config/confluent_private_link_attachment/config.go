package confluent_private_link_attachment

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_private_link_attachment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true
		r.Kind = "PrivateLinkAttachment"

		// Allows us to reference environment by the metadata.name instead of the externally generated (random) name by Confluent.
		r.References["environment.id"] = config.Reference{
			Type: "github.com/stakater/provider-confluent/apis/confluent/v1alpha1.Environment",
		}
	})
}
