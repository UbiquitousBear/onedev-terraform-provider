package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ubiquitousbear/onedev-api"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEDEV_TOKEN", nil),
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEDEV_ADDRESS", nil),
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEDEV_USER", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"onedev_project": resourceOnedevProject(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		address := d.Get("address").(string)
		user := d.Get("user").(string)
		token := d.Get("token").(string)
		return onedev_api.NewClient(address, user, token), nil
	}
}
