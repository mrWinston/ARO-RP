package namespace

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"strings"
)

// IsOpenShift returns true if ns is an openshift managed namespace, including the default namespace
func IsOpenShift(ns string) bool {
	return ns == "" ||
		ns == "default" ||
		ns == "openshift" ||
		strings.HasPrefix(ns, "kube-") ||
		strings.HasPrefix(ns, "openshift-")
}

// IsOpenShiftSystemNamespace returns true if ns is an openshift managed namespace, without the default namespace
func IsOpenShiftSystemNamespace(ns string) bool {
	return ns == "" ||
		ns == "openshift" ||
		strings.HasPrefix(ns, "kube-") ||
		strings.HasPrefix(ns, "openshift-")
}

// FilteredOpenShiftNamespace returns true if ns is a namespace in the defined hardcoded map
func FilteredOpenShiftNamespace(ns string) bool {
	filteredNamespaces := map[string]struct{}{
		"openshift":                                        {},
		"openshift-apiserver":                              {},
		"openshift-apiserver-operator":                     {},
		"openshift-authentication":                         {},
		"openshift-authentication-operator":                {},
		"openshift-cloud-controller-manager":               {},
		"openshift-cloud-controller-manager-operator":      {},
		"openshift-cloud-credential-operator":              {},
		"openshift-cluster-csi-drivers":                    {},
		"openshift-cluster-machine-approver":               {},
		"openshift-cluster-node-tuning-operator":           {},
		"openshift-cluster-samples-operator":               {},
		"openshift-cluster-storage-operator":               {},
		"openshift-config":                                 {},
		"openshift-config-managed":                         {},
		"openshift-config-operator":                        {},
		"openshift-console":                                {},
		"openshift-console-operator":                       {},
		"openshift-console-user-settings":                  {},
		"openshift-controller-manager":                     {},
		"openshift-controller-manager-operator":            {},
		"openshift-dns":                                    {},
		"openshift-dns-operator":                           {},
		"openshift-etcd":                                   {},
		"openshift-etcd-operator":                          {},
		"openshift-host-network":                           {},
		"openshift-image-registry":                         {},
		"openshift-ingress":                                {},
		"openshift-ingress-canary":                         {},
		"openshift-ingress-operator":                       {},
		"openshift-insights":                               {},
		"openshift-kni-infra":                              {},
		"openshift-kube-apiserver":                         {},
		"openshift-kube-apiserver-operator":                {},
		"openshift-kube-controller-manager":                {},
		"openshift-kube-controller-manager-operator":       {},
		"openshift-kube-scheduler":                         {},
		"openshift-kube-scheduler-operator":                {},
		"openshift-kube-storage-version-migrator":          {},
		"openshift-kube-storage-version-migrator-operator": {},
		"openshift-machine-api":                            {},
		"openshift-machine-config-operator":                {},
		"openshift-marketplace":                            {},
		"openshift-monitoring":                             {},
		"openshift-multus":                                 {},
		"openshift-network-diagnostics":                    {},
		"openshift-network-operator":                       {},
		"openshift-oauth-apiserver":                        {},
		"openshift-openstack-infra":                        {},
		"openshift-operator-lifecycle-manager":             {},
		"openshift-operators":                              {},
		"openshift-ovirt-infra":                            {},
		"openshift-sdn":                                    {},
		"openshift-service-ca":                             {},
		"openshift-service-ca-operator":                    {},
		"openshift-user-workload-monitoring":               {},
		"openshift-vsphere-infra":                          {},
		"openshift-azure-operator":                         {},
		"openshift-managed-upgrade-operator":               {},
	}
	_, ok := filteredNamespaces[ns]
	return ok
}
