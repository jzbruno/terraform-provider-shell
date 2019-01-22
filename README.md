# Terraform Shell Provider

## Overview

This is a basic shell provider that includes a data source to allow capturing output of commands
for use with other Terraform resourcees.

In it's current state this provider is just a data source and should not be used to run commands
that modify resources. A future addition of a resource should be used for that use case.

## Install

1. The provider can either be built or downloaded from GitHub. Get the provider.

    * To build the provider

        ```bash
        git clone git@github.com:jzbruno/terraform-provider-shell.git
        cd terraform-provider-shell/
        go get
        go build
        ```
        &NewLine;

    * To download the provider

        ```bash
        curl -sL https://github.com/jzbruno/terraform-provider-shell/releases/download/v0.1.0-alpha/terraform-provider-shell -o terraform-provider-shell
        ```
        &NewLine;

2. Install the provider.

    ```bash
    chmod +x terraform-provider-shell
    mkdir -p ~/.terraform.d/plugins/
    cp terraform-provider-shell ~/.terraform.d/plugins/
    ```
    &NewLine;

    For more information about how Terraform discovers plugins, see [Terraform Plugin Discovery](https://www.terraform.io/docs/extend/how-terraform-works.html#discovery)

## Example

The following example will store your public IP allowing it to be used in another resource

```hcl
data "shell_command" "my_public_ip" {
  command = "curl -s ifconfig.co"
}
```
&NewLine;

To use the output reference the data source in another Terraform resource. If the data source changes 
it will cause an update to the Terraform resource referencing it.

```hcl
resource "aws_security_group" "allow_my_public_ip" {
  name        = "allow_my_public_ip"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["${data.shell_command.my_public_ip.output}/32"]
  }
}
```
&NewLine;