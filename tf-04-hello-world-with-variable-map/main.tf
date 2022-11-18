terraform {
  required_version = "~> 1.3.0"
}

variable "your_name" {
  type        = string
  nullable    = false
  description = "Tell us our name to say hello to you"
}

variable "rock_band" {
  type        = map(string)
  default     = {}
  nullable    = false
  description = "Components of your Rock band"

  validation {
    condition     = length(var.rock_band) > 2
    error_message = "Your rock band has more than 2 members"
  }
}

output "count_rock_band" {
  value = "Hello ${var.your_name}, there are ${length(var.rock_band)} members in your rock band."
}

output "rock_band" {
  value = var.rock_band
}
