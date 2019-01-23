package shell

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rs/xid"
)

const (
	goosWindows = "windows"
)

func dataCommand() *schema.Resource {
	return &schema.Resource{
		Read: dataCommandRead,
		Schema: map[string]*schema.Schema{
			"command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shell": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"trim_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"output": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataCommandRead(d *schema.ResourceData, m interface{}) error {
	command, args, err := commandParse(d)
	if err != nil {
		return fmt.Errorf("error parsing command: %s", err)
	}

	output, err := commandRun(command, args, d.Get("trim_whitespace").(bool))
	if err != nil {
		return fmt.Errorf("error running command: %s", err)
	}

	d.SetId(xid.New().String())
	d.Set("output", output)

	return nil
}
