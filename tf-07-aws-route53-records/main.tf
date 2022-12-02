data "aws_route53_zone" "this" {
    count = var.create && var.zone_name != null ? 1 : 0

    name         = var.zone_name
    private_zone = var.private_zone
}

resource "aws_route53_record" "this" {
  for_each = { for k, v in local.recordsets : k => v if var.create && var.zone_name != null }

  zone_id = data.aws_route53_zone.this[0].zone_id

  name    = each.value.name != "" ? (lookup(each.value, "full_name_override", false) ? each.value.name : "${each.value.name}.${data.aws_route53_zone.this[0].name}") : data.aws_route53_zone.this[0].name
  type    = each.value.type
  ttl     = lookup(each.value, "ttl", null)
  records = try(each.value.records, null)
  #  set_identifier                   = lookup(each.value, "set_identifier", null)
  #  health_check_id                  = lookup(each.value, "health_check_id", null)
  #  multivalue_answer_routing_policy = lookup(each.value, "multivalue_answer_routing_policy", null)
  allow_overwrite = lookup(each.value, "allow_overwrite", false)

  #  dynamic "alias" {
  #    for_each = length(keys(lookup(each.value, "alias", {}))) == 0 ? [] : [true]
  #
  #    content {
  #      name                   = each.value.alias.name
  #      zone_id                = try(each.value.alias.zone_id, data.aws_route53_zone.this[0].zone_id)
  #      evaluate_target_health = lookup(each.value.alias, "evaluate_target_health", false)
  #    }
  #  }
  #
  #  dynamic "failover_routing_policy" {
  #    for_each = length(keys(lookup(each.value, "failover_routing_policy", {}))) == 0 ? [] : [true]
  #
  #    content {
  #      type = each.value.failover_routing_policy.type
  #    }
  #  }
  #
  #  dynamic "weighted_routing_policy" {
  #    for_each = length(keys(lookup(each.value, "weighted_routing_policy", {}))) == 0 ? [] : [true]
  #
  #    content {
  #      weight = each.value.weighted_routing_policy.weight
  #    }
  #  }
  #
  #  dynamic "geolocation_routing_policy" {
  #    for_each = length(keys(lookup(each.value, "geolocation_routing_policy", {}))) == 0 ? [] : [true]
  #
  #    content {
  #      continent   = lookup(each.value.geolocation_routing_policy, "continent", null)
  #      country     = lookup(each.value.geolocation_routing_policy, "country", null)
  #      subdivision = lookup(each.value.geolocation_routing_policy, "subdivision", null)
  #    }
  #  }
}
