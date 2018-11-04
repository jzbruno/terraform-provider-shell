package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jzbruno/terraform-provider-shell/shell"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return shell.Provider()
		},
	})
}
