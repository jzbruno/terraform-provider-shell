package shell

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rs/xid"
)

func resourceCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceCommandCreate,
		Read:   resourceCommandRead,
		Update: resourceCommandUpdate,
		Delete: resourceCommandDelete,

		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"shell": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"trim_whitespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCommandCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(xid.New().String())
	return resourceCommandRead(d, m)
}

func resourceCommandRead(d *schema.ResourceData, m interface{}) error {
	command, args, err := commandParse(d)
	if err != nil {
		return fmt.Errorf("error parsing command: %s", err)
	}

	output, err := commandRun(command, args, d.Get("trim_whitespace").(bool))
	if err != nil {
		return fmt.Errorf("error running command: %s", err)
	}

	d.Set("output", output)

	return nil
}

func resourceCommandUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCommandRead(d, m)
}

func resourceCommandDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
