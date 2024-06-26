package managed_instance_aws_integrations

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

type BalancerType string

const (
	// - ROUTE53 -------------------------
	IntegrationRoute53 commons.FieldName = "integration_route53"
	Domains            commons.FieldName = "domains"
	HostedZoneId       commons.FieldName = "hosted_zone_id"
	SpotinstAcctID     commons.FieldName = "spotinst_acct_id"
	RecordSetType      commons.FieldName = "record_set_type"
	RecordSets         commons.FieldName = "record_sets"
	UsePublicIP        commons.FieldName = "use_public_ip"
	UsePublicDNS       commons.FieldName = "use_public_dns"
	Route53Name        commons.FieldName = "name"
	// -----------------------------------

	LoadBalancers    commons.FieldName = "load_balancers"
	Arn              commons.FieldName = "arn"
	LoadBalancerName commons.FieldName = "name"
	Type             commons.FieldName = "type"
)
