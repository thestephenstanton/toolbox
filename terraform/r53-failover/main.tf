resource "aws_route53_zone" "doomaxe" {
  name = "doomaxe.com"
}

module "doomaxe_failover" {
  source                         = "./failover"
  hosted_zone_zone_id            = "${aws_route53_zone.doomaxe.zone_id}"
  unambiguous_domain_name        = "${aws_route53_zone.doomaxe.name}"
  subdomain                      = ""
  primary_alias_target           = "${module.use1_failover.record_set_name}"
  primary_alias_hosted_zone_id   = "${module.use1_failover.record_set_zone_id}"
  secondary_alias_target         = "${module.use2_failover.record_set_name}"
  secondary_alias_hosted_zone_id = "${module.use2_failover.record_set_zone_id}"
}

module "use1_failover" {
  source                         = "./failover"
  hosted_zone_zone_id            = "${aws_route53_zone.doomaxe.zone_id}"
  unambiguous_domain_name        = "${aws_route53_zone.doomaxe.name}"
  subdomain                      = "use1"
  primary_alias_target           = "dualstack.hyperfang-primary-lb-1795342106.us-east-1.elb.amazonaws.com."
  primary_alias_hosted_zone_id   = "Z35SXDOTRQ7X7K"
  secondary_alias_target         = "dualstack.hyperfang-failover-lb-673098672.us-east-1.elb.amazonaws.com."
  secondary_alias_hosted_zone_id = "Z35SXDOTRQ7X7K"
}

module "use2_failover" {
  source                         = "./failover"
  hosted_zone_zone_id            = "${aws_route53_zone.doomaxe.zone_id}"
  unambiguous_domain_name        = "${aws_route53_zone.doomaxe.name}"
  subdomain                      = "use2"
  primary_alias_target           = "dualstack.hyperfang-primary-lb-2047762972.us-east-2.elb.amazonaws.com."
  primary_alias_hosted_zone_id   = "Z3AADJGX6KTTL2"
  secondary_alias_target         = "dualstack.hyperfang-failover-lb-1423325208.us-east-2.elb.amazonaws.com."
  secondary_alias_hosted_zone_id = "Z3AADJGX6KTTL2"
}
