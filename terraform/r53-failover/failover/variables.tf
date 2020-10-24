variable "hosted_zone_zone_id" {
  type        = "string"
  description = "The ID of the hosted zone to contain this record."
}

variable "unambiguous_domain_name" {
  type        = "string"
  description = "Fully-qualified domain name with . at end. Example: 'google.com.'."
}

variable "subdomain" {
  type        = "string"
  description = "Subdomain of record you want. Example: 'foo' to get 'foo.bar.com'. Don't specify a subdomain to apply failover logic to the root hosted zone's record set."
}

variable "primary_alias_target" {
  type        = "string"
  description = "Fully qualified domain name (A Record) for AWS resource you want for primary routing."
}

variable "primary_alias_hosted_zone_id" {
  type        = "string"
  description = "Hosted zone id for primary_alias_target."
}

variable "secondary_alias_target" {
  type        = "string"
  description = "Fully qualified domain name (A Record) for AWS resource you want for secondary routing."
}

variable "secondary_alias_hosted_zone_id" {
  type        = "string"
  description = "Hosted zone id for secondary_alias_target."
}
