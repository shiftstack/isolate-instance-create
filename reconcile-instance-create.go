package main

import (
	"fmt"

	"gopkg.in/yaml.v2"

	machinev1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/openstack/clients"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/openstack"
	providerClient "sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/openstack/machine"

	openstackconfigv1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/apis/openstackproviderconfig/v1alpha1"
)

const (
	machineSpec = `
apiVersion: machine.openshift.io/v1beta1
kind: MachineSet
metadata:
  annotations:
	machine.openshift.io/memoryMb: "16384"
	machine.openshift.io/vCPU: "8"
  creationTimestamp: "2021-03-17T01:47:53Z"
  generation: 6
  labels:
	machine.openshift.io/cluster-api-cluster: emilio-tqf47
	machine.openshift.io/cluster-api-machine-role: worker
	machine.openshift.io/cluster-api-machine-type: worker
  managedFields:
  - apiVersion: machine.openshift.io/v1beta1
	fieldsType: FieldsV1
	fieldsV1:
	  f:metadata:
		f:labels:
		  .: {}
		  f:machine.openshift.io/cluster-api-cluster: {}
		  f:machine.openshift.io/cluster-api-machine-role: {}
		  f:machine.openshift.io/cluster-api-machine-type: {}
	  f:spec:
		.: {}
		f:selector:
		  .: {}
		  f:matchLabels:
			.: {}
			f:machine.openshift.io/cluster-api-cluster: {}
			f:machine.openshift.io/cluster-api-machineset: {}
		f:template:
		  .: {}
		  f:metadata:
			.: {}
			f:labels:
			  .: {}
			  f:machine.openshift.io/cluster-api-cluster: {}
			  f:machine.openshift.io/cluster-api-machine-role: {}
			  f:machine.openshift.io/cluster-api-machine-type: {}
			  f:machine.openshift.io/cluster-api-machineset: {}
		  f:spec:
			.: {}
			f:metadata: {}
			f:providerSpec:
			  .: {}
			  f:value:
				.: {}
				f:apiVersion: {}
				f:cloudName: {}
				f:cloudsSecret: {}
				f:flavor: {}
				f:image: {}
				f:kind: {}
				f:metadata: {}
				f:securityGroups: {}
				f:serverMetadata: {}
				f:tags: {}
				f:trunk: {}
				f:userDataSecret: {}
	  f:status:
		.: {}
		f:observedGeneration: {}
	manager: cluster-bootstrap
	operation: Update
	time: "2021-03-17T01:47:53Z"
  - apiVersion: machine.openshift.io/v1beta1
	fieldsType: FieldsV1
	fieldsV1:
	  f:metadata:
		f:annotations:
		  .: {}
		  f:machine.openshift.io/memoryMb: {}
		  f:machine.openshift.io/vCPU: {}
	manager: machine-controller-manager
	operation: Update
	time: "2021-03-17T01:53:11Z"
  - apiVersion: machine.openshift.io/v1beta1
	fieldsType: FieldsV1
	fieldsV1:
	  f:spec:
		f:replicas: {}
		f:template:
		  f:spec:
			f:providerSpec:
			  f:value:
				f:networks: {}
	manager: kubectl-edit
	operation: Update
	time: "2021-03-18T15:23:58Z"
  - apiVersion: machine.openshift.io/v1beta1
	fieldsType: FieldsV1
	fieldsV1:
	  f:status:
		f:availableReplicas: {}
		f:fullyLabeledReplicas: {}
		f:observedGeneration: {}
		f:readyReplicas: {}
		f:replicas: {}
	manager: machineset-controller
	operation: Update
	time: "2021-03-18T15:29:45Z"
  name: emilio-tqf47-worker-sriov
  namespace: openshift-machine-api
  resourceVersion: "694559"
  uid: bc377d66-6fc6-40ff-8431-12c6147097c5
spec:
  replicas: 2
  selector:
	matchLabels:
	  machine.openshift.io/cluster-api-cluster: emilio-tqf47
	  machine.openshift.io/cluster-api-machineset: emilio-tqf47-worker-sriov
  template:
	metadata:
	  labels:
		machine.openshift.io/cluster-api-cluster: emilio-tqf47
		machine.openshift.io/cluster-api-machine-role: worker
		machine.openshift.io/cluster-api-machine-type: worker
		machine.openshift.io/cluster-api-machineset: emilio-tqf47-worker-sriov
	spec:
	  metadata: {}
	  providerSpec:
		value:
		  apiVersion: openstackproviderconfig.openshift.io/v1alpha1
		  cloudName: openstack
		  cloudsSecret:
			name: openstack-cloud-credentials
			namespace: openshift-machine-api
		  flavor: ci.m1.xlarge
		  image: rhcos-4.8
		  kind: OpenstackProviderSpec
		  metadata:
			creationTimestamp: null
		  networks:
		  - filter: {}
			subnets:
			- count: 2
			  portSecurity: false
			  filter:
				name: emilio-tqf47-nodes
				tags: openshiftClusterID=emilio-tqf47
			  portTags:
			  - sriov
		  securityGroups:
		  - filter: {}
			name: emilio-tqf47-worker
		  serverMetadata:
			Name: emilio-tqf47-worker
			openshiftClusterID: emilio-tqf47
		  tags:
		  - openshiftClusterID=emilio-tqf47
		  trunk: false
		  userDataSecret:
			name: worker-user-data
status:
  availableReplicas: 2
  fullyLabeledReplicas: 2
  observedGeneration: 6
  readyReplicas: 2
  replicas: 2
`
)

func main() {
	machine := &machinev1.Machine{}
	yaml.Unmarshal([]byte(machineSpec), machine)
	providerSpec, err := openstackconfigv1.MachineSpecFromProviderSpec(machine.Spec.ProviderSpec)
	if err != nil {
		panic(fmt.Errorf("Could not parse providerSpec from MachineSpec: %v", err))
	}

	actuatorParams := openstack.ActuatorParams{

	}
	oc, err := providerClient.NewActuator(actuatorParams)
	if err != nil {
		panic(fmt.Errorf("Failed to create new actuator: %v", err))
	}

	kubeClient := providerClient.OpenstackClient.

	// TODO: replace newInstanceServiceFromMachine with a function that get OS_CLOUD from local env
	machineService, err := clients.NewInstanceServiceFromMachine(kubeClient, machine)
	if err != nil {
		panic(fmt.Errorf("Error creating new instance service: %v", err))
	}

	var userData, clusterName, keyName string
	clusterName = "test"

	instance, err := machineService.InstanceCreate(clusterName, machine.Name, &openstackconfigv1.OpenstackClusterProviderSpec{}, providerSpec, userData, keyName, providerClient.OpenstackClient.params.configClient)
	if err != nil {
		panic(fmt.Errorf("Unable to create OpenStack instance from Machine %s: %v", machine.Name, err))
	}

	fmt.Printf("Instance %s Created", instance.ID)
}
