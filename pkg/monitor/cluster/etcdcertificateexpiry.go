package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"math"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	utilcert "github.com/Azure/ARO-RP/pkg/util/cert"
	utilpem "github.com/Azure/ARO-RP/pkg/util/pem"
	"github.com/Azure/ARO-RP/pkg/util/version"
)

func (mon *Monitor) emitEtcdCertificateExpiry(ctx context.Context) error {
	cv, err := mon.getClusterVersion(ctx)
	if err != nil {
		return err
	}
	v, err := version.ParseVersion(actualVersion(cv))
	if err != nil {
		return err
	}
	// ETCD ceritificates are autorotated by the operator when close to expiry for cluster running 4.9+
	if v.Lt(version.NewVersion(4, 9)) {
		return nil
	}

	secretList, err := mon.cli.CoreV1().Secrets("openshift-etcd").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("type=%s", corev1.SecretTypeTLS)})
	if err != nil {
		return err
	}

	certNearExpiry := false
	minDaysUntilExpiration := math.MaxInt
	for _, secret := range secretList.Items {
		if strings.Contains(secret.ObjectMeta.Name, "etcd-peer") || strings.Contains(secret.ObjectMeta.Name, "etcd-serving") {
			_, certs, err := utilpem.Parse(secret.Data[corev1.TLSCertKey])
			if err != nil {
				return err
			}
			if !utilcert.IsLessThanMinimumDuration(certs[0], utilcert.DefaultMinDurationPercent) {
				certNearExpiry = true
				minDaysUntilExpiration = min(utilcert.DaysUntilExpiration(certs[0]), minDaysUntilExpiration)
			}
		}
	}

	if certNearExpiry {
		mon.emitGauge("certificate.expirationdate", 1, map[string]string{
			"daysUntilExpiration": fmt.Sprintf("%d", minDaysUntilExpiration),
			"namespace":           "openshift-etcd",
			"name":                "openshift-etcd-certificate",
		})
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
