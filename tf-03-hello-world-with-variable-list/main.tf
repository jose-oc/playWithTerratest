terraform {
  required_version = "~> 1.3.0"
}

variable "your_name" {
  type        = string
  nullable    = false
  description = "Tell us our name to say hello to you"
}

variable "friends" {
  type        = list(string)
  default     = []
  nullable    = false
  description = "Tell us your friends' names"

  validation {
    condition     = length(var.friends) > 1
    error_message = "You have to name at least 2 friends."
  }
}

locals {
  greeting_friends = formatlist("Hello, %s!", var.friends)
}

output "count_friends" {
  value = "Hello ${var.your_name}, you have ${length(var.friends)} friends."
}

output "greeting_friends" {
  value = local.greeting_friends
}