package provider

import (
	"context"
	c "github.com/cloudposse/atmos/pkg/convert"
	l "github.com/cloudposse/terraform-provider-atmos/internal/label"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLabel() *schema.Resource {
	return &schema.Resource{
		Description: "The `label` data source accepts a context " +
			"and returns an ID",

		ReadContext: dataSourceLabelRead,

		Schema: map[string]*schema.Schema{
			"namespace": {
				Description: "Namespace.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tenant": {
				Description: "Tenant.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"environment": {
				Description: "Environment.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"stage": {
				Description: "Stage.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delimiter": {
				Description: "Delimiter.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"output": {
				Description: "Label.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceLabelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	namespace := d.Get("namespace").(string)
	tenant := d.Get("tenant").(string)
	environment := d.Get("environment").(string)
	stage := d.Get("stage").(string)
	name := d.Get("name").(string)
	delimiter := d.Get("delimiter").(string)

	label, err := l.CreateLabel(namespace, tenant, environment, stage, name, delimiter)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("output", string(label))
	if err != nil {
		return diag.FromErr(err)
	}

	id := c.MakeId([]byte(label))
	d.SetId(id)

	return nil
}
