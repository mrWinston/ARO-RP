package cluster

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	configv1 "github.com/openshift/api/config/v1"
	configfake "github.com/openshift/client-go/config/clientset/versioned/fake"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	mock_metrics "github.com/Azure/ARO-RP/pkg/util/mocks/metrics"
	utiltls "github.com/Azure/ARO-RP/pkg/util/tls"
)

func TestEtcdCertificateExpiry(t *testing.T) {
	ctx := context.Background()
	expiration := time.Now().Add(time.Microsecond * 60)
	_, cert, err := utiltls.GenerateTestKeyAndCertificate("etcd-cert", nil, nil, false, false, tweakTemplateFn(expiration))
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range []struct {
		name                   string
		configcli              *configfake.Clientset
		cli                    *fake.Clientset
		minDaysUntilExpiration int
	}{
		{
			name: "emit etcd certificate expiry",
			configcli: configfake.NewSimpleClientset(
				&configv1.ClusterVersion{
					ObjectMeta: metav1.ObjectMeta{
						Name: "version",
					},
					Status: configv1.ClusterVersionStatus{
						History: []configv1.UpdateHistory{
							{
								State:   configv1.CompletedUpdate,
								Version: "4.8.1",
							},
						},
					},
				},
			),
			cli: fake.NewSimpleClientset(
				&corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "etcd-peer-master-0",
						Namespace: "openshift-etcd",
					},
					Data: map[string][]byte{
						corev1.TLSCertKey: pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert[0].Raw}),
					},
					Type: corev1.SecretTypeTLS,
				},
			),
			minDaysUntilExpiration: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			m := mock_metrics.NewMockEmitter(controller)
			mon := &Monitor{
				cli:       tt.cli,
				configcli: tt.configcli,
				m:         m,
			}

			m.EXPECT().EmitGauge("certificate.expirationdate", int64(1), map[string]string{
				"daysUntilExpiration": fmt.Sprintf("%d", tt.minDaysUntilExpiration),
				"namespace":           "openshift-etcd",
				"name":                "openshift-etcd-certificate",
			})

			err = mon.emitEtcdCertificateExpiry(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func tweakTemplateFn(expiration time.Time) func(*x509.Certificate) {
	return func(template *x509.Certificate) {
		template.NotAfter = expiration
	}
}
