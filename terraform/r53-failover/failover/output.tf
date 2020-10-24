output "record_set_name" {
  value = "${aws_route53_record.primary.name}"
}

output "record_set_zone_id" {
  value = "${aws_route53_record.primary.zone_id}"
}
