package shell

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"shell_command": dataCommand(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"shell_command": resourceCommand(),
		},
	}
}
