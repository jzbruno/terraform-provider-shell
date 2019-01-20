provider "aws" {
  profile = "jzbruno-terraform"
  region  = "us-east-1"
}

data "shell_command" "my_public_ip" {
  command = "curl -s ifconfig.co"
}

resource "aws_security_group" "allow_my_public_ip" {
  name        = "allow_my_public_ip"
  description = "Example usage of the terraform-shell provider to get and allow my public ip."

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["${data.shell_command.my_public_ip.output}/32"]
  }

  tags = {
    Name = "allow_my_public_ip"
  }
}
