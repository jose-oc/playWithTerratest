data "aws_route53_zone" "this" {
  name         = var.zone_name
  private_zone = var.private_zone
}

resource "aws_route53_record" "this" {
  name    = "${var.record_prefix_name}.${var.zone_name}"
  type    = "A"
  zone_id = data.aws_route53_zone.this.zone_id
  records = [aws_instance.webserver.public_ip]
  ttl     = var.ttl
}