module github.com/shiftstack/isolate-instance-create

go 1.15

require (
	github.com/openshift/machine-api-operator v0.2.1-0.20210104142355-8e6ae0acdfcf
	gopkg.in/yaml.v2 v2.4.0
	sigs.k8s.io/cluster-api-provider-aws v0.0.0 // indirect
	sigs.k8s.io/cluster-api-provider-azure v0.0.0 // indirect
	sigs.k8s.io/cluster-api-provider-openstack v0.0.0
)

replace (
	sigs.k8s.io/cluster-api-provider-aws => github.com/openshift/cluster-api-provider-aws v0.2.1-0.20210121023454-5ffc5f422a80
	sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20201016155852-4090a6970205
	sigs.k8s.io/cluster-api-provider-openstack => github.com/openshift/cluster-api-provider-openstack v0.0.0-20201116051540-155384b859c5
)
