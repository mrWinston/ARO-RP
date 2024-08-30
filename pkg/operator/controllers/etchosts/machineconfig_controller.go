package etchosts

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"regexp"

	mcv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/Azure/ARO-RP/pkg/operator"
	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	"github.com/Azure/ARO-RP/pkg/operator/controllers/base"
	"github.com/Azure/ARO-RP/pkg/operator/predicates"
	"github.com/Azure/ARO-RP/pkg/util/dynamichelper"
)

const (
	ControllerName = "EtcHostsMachineConfig"
)

type EtcHostsMachineConfigReconciler struct {
	base.AROController

	dh dynamichelper.Interface
}

var etcHostsRegex = regexp.MustCompile("^99-(.*)-aro-etc-hosts-gateway-domains$")

func NewReconciler(log *logrus.Entry, client client.Client, dh dynamichelper.Interface) *EtcHostsMachineConfigReconciler {
	return &EtcHostsMachineConfigReconciler{
		AROController: base.AROController{
			Log:    log,
			Client: client,
			Name:   ControllerName,
		},
		dh: dh,
	}
}

// Reconcile watches ARO EtcHosts MachineConfig objects, and if any changes, reconciles it
func (r *EtcHostsMachineConfigReconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	r.Log.Debugf("reconcile MachineConfig openshift-machine-api/%s", request.Name)

	instance, err := r.GetCluster(ctx)
	if err != nil {
		return reconcile.Result{}, err
	}

	if !instance.Spec.OperatorFlags.GetSimpleBoolean(operator.EtcHostsEnabled) {
		r.Log.Debug("controller is disabled")
		return reconcile.Result{}, nil
	}

	r.Log.Debug("running")
	mcp := &mcv1.MachineConfigPool{}
	// Make sure we are reconciling against etchosts machine config
	m := etcHostsRegex.FindStringSubmatch(request.Name)
	r.Log.Debugf("MachineConfig openshift-machine-api/%s", request.Name)
	if m == nil {
		return reconcile.Result{}, nil
	}
	role := m[1]

	err = r.Client.Get(ctx, types.NamespacedName{Name: role}, mcp)
	if kerrors.IsNotFound(err) {
		r.ClearDegraded(ctx)
		return reconcile.Result{}, nil
	}
	if err != nil {
		r.Log.Error(err)
		r.SetDegraded(ctx, err)
		return reconcile.Result{}, err
	}
	if mcp.GetDeletionTimestamp() != nil {
		return reconcile.Result{}, nil
	}

	err = reconcileMachineConfigs(ctx, instance, role, r.dh, *mcp)
	if err != nil {
		r.Log.Error(err)
		r.SetDegraded(ctx, err)
		return reconcile.Result{}, err
	}

	r.ClearConditions(ctx)
	return reconcile.Result{}, nil
}

// SetupWithManager setup our mananger to watch for changes to MCP and ARO Cluster obj
func (r *EtcHostsMachineConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.Log.Info("starting etchosts-machine-config controller")

	etcHostsBuilder := ctrl.NewControllerManagedBy(mgr).
		For(&mcv1.MachineConfig{}).
		Watches(&source.Kind{Type: &mcv1.MachineConfigPool{}},
			&handler.EnqueueRequestForObject{},
			builder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Watches(&source.Kind{Type: &arov1alpha1.Cluster{}},
			&handler.EnqueueRequestForObject{},
			builder.WithPredicates(predicate.And(predicates.AROCluster, predicate.GenerationChangedPredicate{})))

	return etcHostsBuilder.
		WithEventFilter(predicate.Or(predicate.GenerationChangedPredicate{}, predicate.AnnotationChangedPredicate{}, predicate.LabelChangedPredicate{})).
		Named(ControllerName).
		Complete(r)
}

func reconcileMachineConfigs(ctx context.Context, instance *arov1alpha1.Cluster, role string, dh dynamichelper.Interface, mcps ...mcv1.MachineConfigPool) error {
	var resources []kruntime.Object
	for _, mcp := range mcps {
		resource, err := EtcHostsMachineConfig(instance.Spec.Domain, instance.Spec.APIIntIP, instance.Spec.GatewayDomains, instance.Spec.GatewayPrivateEndpointIP, role)
		if err != nil {
			return err
		}

		err = dynamichelper.SetControllerReferences([]kruntime.Object{resource}, &mcp)
		if err != nil {
			return err
		}

		resources = append(resources, resource)
	}

	err := dynamichelper.Prepare(resources)
	if err != nil {
		return err
	}

	return dh.Ensure(ctx, resources...)
}