package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support Markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

// New creates a new provider and returns a *schema.Provider
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"atmos_stack_config":              dataSourceStackConfig(),
				"atmos_spacelift_stack_config":    dataSourceSpaceliftStackConfig(),
				"atmos_component_config":          dataSourceComponentConfig(),
				"atmos_aws_eks_update_kubeconfig": dataSourceAwsEksUpdateKubeconfig(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return nil, nil
	}
}
