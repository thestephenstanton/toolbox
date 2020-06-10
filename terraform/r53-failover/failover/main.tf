locals {
  record_set_name       = "${var.subdomain == "" ? "${var.unambiguous_domain_name}" : "${var.subdomain}.${var.unambiguous_domain_name}"}"
  set_identifier_prefix = "${var.subdomain == "" ? "" : "${var.subdomain}-"}"
}

resource "aws_route53_record" "primary" {
  zone_id = "${var.hosted_zone_zone_id}"
  name    = "${local.record_set_name}"
  type    = "A"

  alias {
    name                   = "${var.primary_alias_target}"
    zone_id                = "${var.primary_alias_hosted_zone_id}"
    evaluate_target_health = true
  }

  failover_routing_policy {
    type = "PRIMARY"
  }

  set_identifier = "${local.set_identifier_prefix}primary"
}

resource "aws_route53_record" "secondary" {
  zone_id = "${var.hosted_zone_zone_id}"
  name    = "${local.record_set_name}"
  type    = "A"

  alias {
    name                   = "${var.secondary_alias_target}"
    zone_id                = "${var.secondary_alias_hosted_zone_id}"
    evaluate_target_health = true
  }

  failover_routing_policy {
    type = "SECONDARY"
  }

  set_identifier = "${local.set_identifier_prefix}secondary"
}
