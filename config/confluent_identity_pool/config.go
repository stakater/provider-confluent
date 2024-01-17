package confluent_identity_pool

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_identity_pool", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true
		r.Kind = "IdentityPool"

		r.References["identity_provider.id"] = config.Reference{
			Type: "github.com/stakater/provider-confluent/apis/confluent/v1alpha1.IdentityProvider",
		}
	})
}
