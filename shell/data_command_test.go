package shell

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceCommand(t *testing.T) {
	resourceName := "data.shell_command.test"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"shell": Provider().(*schema.Provider),
		},
		Steps: []resource.TestStep{
			{
				Config: testDataShellCommandConfig,
				Check: checkDataSourceAttributes(resourceName),
			},
		},
	})
}

func checkDataSourceAttributes(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		// Ensure the state is populated with all the computed attributes defined by the resource schema.
		for k, v := range dataCommand().Schema {
			if !v.Computed {
				continue
			}

			if _, ok := rs.Primary.Attributes[k]; !ok {
				return fmt.Errorf("state missing attribute %s", k)
			}
		}

		return nil
	}
}

const testDataShellCommandConfig = `
data "shell_command" "test" {
  command = "curl -s ifconfig.co"
}
`
