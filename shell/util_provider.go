package shell

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"shell_command": dataCommand(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"shell_command": resourceCommand(),
		},
	}
}
