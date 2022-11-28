# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# AWS_ACCESS_KEY_ID
# AWS_SECRET_ACCESS_KEY

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# ---------------------------------------------------------------------------------------------------------------------
variable "aws_account_id" {
  description = "AWS account ID so we are sure we apply this code to the right account"
  type        = string
}

variable "main_vpc_cidr" {
  description = "The CIDR of the main VPC"
  type        = string
}

variable "public_subnet_cidr" {
  description = "The CIDR of public subnet"
  type        = string
}

variable "private_subnet_cidr" {
  description = "The CIDR of the private subnet"
  type        = string
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "aws_region" {
  description = "The AWS region to deploy into"
  type        = string
  default     = "us-west-2"
}

variable "tag_name" {
  description = "A name used to tag the resource"
  type        = string
  default     = "terratest-network-poc"
}
