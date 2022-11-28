provider "aws" {
  allowed_account_ids = [var.aws_account_id]
  default_tags {
    tags = {
      App       = local.app
      Name      = "terratest-poc-webserver"
      BuiltWith = "terraform"
      Goal      = local.goal
    }
  }
  region = var.aws_region
}

provider "tls" {}

# -- Generate Keys to ssh into the EC2 instance

resource "tls_private_key" "tls_key" {
  # Avoid creating certificates with terraform in production as the tfplan holds sensitive information in plain text
  algorithm = "RSA"
  rsa_bits  = 4096
}

# Key
resource "aws_key_pair" "generated_key" {
  key_name = "terratest-poc-webserver"

  public_key = tls_private_key.tls_key.public_key_openssh

  # Store private key :  Generate and save private key(terratest-poc-webserver.pem) in current directory
  provisioner "local-exec" {
    command = <<-EOT
      echo '${tls_private_key.tls_key.private_key_pem}' > terratest-poc-webserver.pem
      chmod 400 terratest-poc-webserver.pem
    EOT
  }
}


data "aws_ami" "ami" {
  most_recent = true

  owners = [
    "amazon",
  ]

  filter {
    name = "name"

    values = [
      "amzn2-ami-kernel-5.10-hvm-*-x86_64*"
    ]
  }
}

# Deploy an EC2 Instance.
resource "aws_instance" "webserver" {
  ami           = data.aws_ami.ami.image_id
  instance_type = "t2.micro"
  # Allow ssh into the EC2 in case it's needed. You can remove it for security
  vpc_security_group_ids = [
    aws_security_group.ssh.id,
    aws_security_group.webserver.id
  ]

  key_name = aws_key_pair.generated_key.key_name

  # When the instance boots, start a web server on port 80 that responds with text
  user_data = <<EOF
#!/bin/bash
# get admin privileges
sudo su

# install httpd (Linux 2 version)
yum update -y
yum install -y httpd.x86_64
systemctl start httpd.service
systemctl enable httpd.service
echo "Hello World from $(hostname -f)" > /var/www/html/index.html
EOF
}

# Allow the instance to receive HTTP(s) requests
resource "aws_security_group" "webserver" {
  # HTTPS
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # HTTP
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Allow ssh into the EC2 in case it's needed.
resource "aws_security_group" "ssh" {
  egress {
    from_port   = 22
    protocol    = "tcp"
    to_port     = 22
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
