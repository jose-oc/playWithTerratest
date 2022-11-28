output "public_ip" {
  value       = aws_instance.webserver.public_ip
  description = "The web server public IP"
}

output "private_dns" {
  value = aws_instance.webserver.private_dns
}

output "private_ip" {
  value = aws_instance.webserver.private_ip
}

output "webserver_dns_url" {
  value = aws_route53_record.this.name
}