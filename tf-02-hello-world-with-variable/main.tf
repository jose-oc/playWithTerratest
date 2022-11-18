terraform {
  required_version = "~> 1.3.0"
}

variable "your_name" {
  type = string
  description = "Tell us our name to say hello to you"
}

output "greetings" {
  value = "Hello, ${var.your_name}!"
}