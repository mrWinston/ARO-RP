package billing

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/env"
)

const (
	tenantIDMSFT = "72f988bf-86f1-41af-91ab-2d7cd011db47"
	tenantIDAME  = "33e01921-4d64-4f8c-a055-5bdaffd5e33d"
)

type Manager interface {
	Ensure(context.Context, *api.OpenShiftClusterDocument, *api.SubscriptionDocument) error
	Delete(context.Context, *api.OpenShiftClusterDocument) error
}

type manager struct {
	billingDB database.Billing
	subDB     database.Subscriptions
	log       *logrus.Entry
}

func NewManager(env env.Interface, billing database.Billing, sub database.Subscriptions, log *logrus.Entry) (Manager, error) {
	return &manager{
		subDB:     sub,
		billingDB: billing,
		log:       log,
	}, nil
}

func (m *manager) Ensure(ctx context.Context, doc *api.OpenShiftClusterDocument, sub *api.SubscriptionDocument) error {
	_, err := m.billingDB.Create(ctx, &api.BillingDocument{
		ID:                        doc.ID,
		Key:                       doc.Key,
		ClusterResourceGroupIDKey: doc.ClusterResourceGroupIDKey,
		InfraID:                   doc.OpenShiftCluster.Properties.InfraID,
		Billing: &api.Billing{
			TenantID: sub.Subscription.Properties.TenantID,
			Location: doc.OpenShiftCluster.Location,
		},
	})
	if err, ok := err.(*cosmosdb.Error); ok &&
		err.StatusCode == http.StatusConflict {
		m.log.Print("billing record already present in DB")
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

func (m *manager) Delete(ctx context.Context, doc *api.OpenShiftClusterDocument) error {
	m.log.Printf("updating billing record with deletion time")
	_, err := m.billingDB.MarkForDeletion(ctx, doc.ID)
	if cosmosdb.IsErrorStatusCode(err, http.StatusNotFound) {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}
