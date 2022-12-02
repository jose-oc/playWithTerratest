variable "create" {
  description = "Whether to create DNS records"
  type        = bool
}

variable "zone_name" {
  description = "Name of DNS zone"
  type        = string
}

variable "private_zone" {
  description = "Whether Route53 zone is private or public"
  type        = bool
}

variable "records" {
  description = "List of objects of DNS records"
  type        = any
}

variable "records_jsonencoded" {
  description = "List of map of DNS records (stored as jsonencoded string, for terragrunt)"
  type        = string
  default     = null
}
