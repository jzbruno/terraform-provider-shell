package shell

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func commandParse(d *schema.ResourceData) (command string, args []string, err error) {
	shell := d.Get("shell").([]interface{})
	if len(shell) > 0 {
		command = shell[0].(string)

		for _, i := range shell[1:] {
			if arg, ok := i.(string); ok {
				args = append(args, arg)
			}
		}
	} else {
		if runtime.GOOS == goosWindows {
			command = "cmd"
			args = []string{"/C"}
		} else {
			command = "/bin/bash"
			args = []string{"-c"}
		}
	}

	actualCommand := d.Get("command").(string)
	if actualCommand == "" {
		err = fmt.Errorf("command must not be an empty string")
		return
	}

	args = append(args, actualCommand)

	return
}

func commandRun(command string, args []string, trim bool) (output string, err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("error running command '%s %v': %s", command, args, err)
		return
	}

	if cmdError := stderr.String(); cmdError != "" {
		err = fmt.Errorf("error running command '%s': %s", args[len(args)-1], cmdError)
		return
	}

	output = stdout.String()
	if trim {
		output = strings.TrimSpace(output)
	}

	return
}
