package confluent_network_link_endpoint

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_network_link_endpoint", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true
		r.Kind = "NetworkLinkEndpoint"

		// Allows us to reference environment by the metadata.name instead of the externally generated (random) name by Confluent.
		r.References["environment.id"] = config.Reference{
			Type: "github.com/stakater/provider-confluent/apis/confluent/v1alpha1.Environment",
		}

		r.References["network.id"] = config.Reference{
			Type: "github.com/stakater/provider-confluent/apis/confluent/v1alpha1.Network",
		}
	})
}
