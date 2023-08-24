package reconcilers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kuadrant/gateway-api-machinery/common"
)

type TargetRefReconciler struct {
	client.Client
}

// ReconcileTargetBackReference adds the policy key in the annotations of the target object
func (r *TargetRefReconciler) ReconcileTargetBackReference(ctx context.Context, policyKey client.ObjectKey, targetNetworkObject client.Object, annotationName string) error {
	logger, _ := logr.FromContext(ctx)

	targetNetworkObjectKey := client.ObjectKeyFromObject(targetNetworkObject)
	targetNetworkObjectKind := targetNetworkObject.GetObjectKind().GroupVersionKind()

	// Reconcile the back reference:
	objAnnotations := common.ReadAnnotationsFromObject(targetNetworkObject)

	if val, ok := objAnnotations[annotationName]; ok {
		if val != policyKey.String() {
			return fmt.Errorf("the %s target %s is already referenced by policy %s", targetNetworkObjectKind, targetNetworkObjectKey, policyKey.String())
		}
	} else {
		objAnnotations[annotationName] = policyKey.String()
		targetNetworkObject.SetAnnotations(objAnnotations)
		err := r.Client.Update(ctx, targetNetworkObject)
		logger.V(1).Info("ReconcileTargetBackReference: update target object", "kind", targetNetworkObjectKind, "name", targetNetworkObjectKey, "err", err)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteTargetBackReference removes the policy key from the annotations of the target object
func (r *TargetRefReconciler) DeleteTargetBackReference(ctx context.Context, _ client.ObjectKey, targetNetworkObject client.Object, annotationName string) error {
	logger, _ := logr.FromContext(ctx)

	targetNetworkObjectKey := client.ObjectKeyFromObject(targetNetworkObject)
	targetNetworkObjectKind := targetNetworkObject.GetObjectKind().GroupVersionKind()

	// Reconcile the back reference:
	objAnnotations := common.ReadAnnotationsFromObject(targetNetworkObject)

	if _, ok := objAnnotations[annotationName]; ok {
		delete(objAnnotations, annotationName)
		targetNetworkObject.SetAnnotations(objAnnotations)
		err := r.Client.Update(ctx, targetNetworkObject)
		logger.V(1).Info("DeleteTargetBackReference: update network resource", "kind", targetNetworkObjectKind, "name", targetNetworkObjectKey, "err", err)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReconcileGatewayPolicyReferences updates in the Gateway resources the annotations that list all the policies
// that directly or indirectly target the gateway, based on a pre-computed gateway diff object
// TODO(@guicassolato): unit test
func (r *TargetRefReconciler) ReconcileGatewayPolicyReferences(ctx context.Context, policy client.Object, gwDiffObj *GatewayDiffs) error {
	logger, _ := logr.FromContext(ctx)

	// delete the policy from the annotations of the gateways no longer targeted by the policy
	for _, gw := range gwDiffObj.GatewaysWithInvalidPolicyRef {
		if gw.DeletePolicy(client.ObjectKeyFromObject(policy)) {
			err := r.Client.Update(ctx, gw.Gateway)
			logger.V(1).Info("ReconcileGatewayPolicyReferences: update gateway", "gateway with invalid policy ref", gw.Key(), "err", err)
			if err != nil {
				return err
			}
		}
	}

	// add the policy to the annotations of the gateways targeted by the policy
	for _, gw := range gwDiffObj.GatewaysMissingPolicyRef {
		if gw.AddPolicy(client.ObjectKeyFromObject(policy)) {
			err := r.Client.Update(ctx, gw.Gateway)
			logger.V(1).Info("ReconcileGatewayPolicyReferences: update gateway", "gateway missing policy ref", gw.Key(), "err", err)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
