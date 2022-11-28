variable "aws_account_id" {
  description = "AWS account ID so we are sure we apply this code to the right account"
  type        = string
}

variable "aws_region" {
  description = "The AWS region to deploy into"
  type        = string
  default     = "us-west-2"
}

variable "zone_name" {
  description = "Name of DNS zone"
  type        = string
  default     = null
}

variable "private_zone" {
  description = "Whether Route53 zone is private or public"
  type        = bool
  default     = false
}

variable "record_prefix_name" {
  description = "Prefix of the record for Route53, it'll be followed by the zone_name"
  type        = string
}

variable "ttl" {
  description = "Route53 record TTL"
  type        = number
  default     = 300
}