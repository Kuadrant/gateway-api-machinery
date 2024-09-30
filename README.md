**⚠️ DEPRECATED**

This repo has been deprecated and will no longer be maintained.
Please refer instead to Kuadrant [**Policy Machinery**](https://github.com/kuadrant/policy-machinery).

----

# Kuadrant Gateway API Machinery

Kuadrant's shared library for implementing Gateway API policy controllers.

## Model

<picture>
  <source media="(prefers-color-scheme: light)" srcset="http://cdn-0.plantuml.com/plantuml/png/ZOx1JeKm44Nt_Of9tMWYF_0MvhrkT4a8EtKnqA4XbfQKQGX_hrHhgb7oTctElNFkUM4C72Sh0lMCpbW2-OXCAsuIS06pbkIfRUl6HwR4mlugSUtjs6zmINHEdyiVN1LS2P7EG1Ndwk531oUOCP3ZXWOXlev0zNpJsKYlES8O3AM8MJEyrvwPz9x9jHD8FUu3eCF-3G8DB-uMdVECF7ft9xD1bOOqg9JaYVzur3MUpvs1z7VzvyvSN7ut3uhgi2ZEv7ISx3i0">
  <img alt="Model" src="http://cdn-0.plantuml.com/plantuml/dpng/ZOx1JeKm44Nt_Of9tMWYF_0MvhrkT4a8EtKnqA4XbfQKQGX_hrHhgb7oTctElNFkUM4C72Sh0lMCpbW2-OXCAsuIS06pbkIfRUl6HwR4mlugSUtjs6zmINHEdyiVN1LS2P7EG1Ndwk531oUOCP3ZXWOXlev0zNpJsKYlES8O3AM8MJEyrvwPz9x9jHD8FUu3eCF-3G8DB-uMdVECF7ft9xD1bOOqg9JaYVzur3MUpvs1z7VzvyvSN7ut3uhgi2ZEv7ISx3i0">
</picture>

<br/>

Kuadrant annotates [Gateway API](https://gateway-api.sigs.k8s.io) resources (`Gateway` and `HTTPRoutes`) with back refences to the custom policies that extend the behavior of the resource.
This allows processing reconciliation events from the policy custom resources themselves to the affected Gateway API resources and the otherway around.

Each kind of policy defines an annotation name. The names of policies that affect a Gateway API resource directly or indirectly referenced by the policies in their `spec.targetRef` fields is stored and kept up to date in the value of that annotation. The value of the annotation is a stringified JSON array of serialized Kubernetes [`types.NamespacedName`](https://pkg.go.dev/k8s.io/apimachinery@v0.27.2/pkg/types#NamespacedName) objects, where each object represents a policy that affects the annotated Gateway API resource.

The policy back reference annotations in the Gateway API resources also allow controlling when a given Gateway API resource has been claimed as a direct `targetRef` of the policy already or not, in cases where no more than one policy of a policy kind can target the same Gateway API resource.

## Library

### Types

**`Referrer` (interface)**<br/>
One which refers to a target object and therefore is referred back from the target object.

**`GatewayWrapper`**<br/>
Wraps a Gateway API `Gateway` resource for a particular `Referrer` implementation.

### Helper functions

**`FetchTargetRefObject`**<br/>
Fetches the target reference object and checks if the status of the resource is valid.

**`ComputeGatewayDiffs`**<br/>
Computes all the differences to reconcile regarding the gateways whose behaviors should/should not be extended by the policy.
These include gateways directly referenced by the policy and gateways indirectly referenced through the policy's target network objects.
- list of gateways to which the policy applies for the first time
- list of gateways to which the policy no longer applies
- list of gateways to which the policy still applies

### Reconciliation functions

Functions to reconcile back references from targeted network objects

**`ReconcileTargetBackReference`**<br/>
Stores the policy key in the annotations of the target object.

**`DeleteTargetBackReference`**<br/>
Removes the policy key from the annotations of the target object.

**`ReconcileGatewayPolicyReferences`**<br/>
Updates in the `Gateway` resources the annotations that list all the policies that directly or indirectly target the gateway, based on a pre-computed gateway diff object.

### Mapping functions

Functions to map Gateway API resource to policies upon reconciliation events trigerred for the Gateway API resources.

| Network resource kind | Mapper constructor function   |
| --------------------- | ----------------------------- |
| `Gateway`             | **`NewGatewayEventMapper`**   |
| `HTTPRoute`           | **`NewHTTPRouteEventMapper`** |

Usage:

```go
func (r *MyPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
  logger := log.NewLogger()

	gatewayEventMapper := mappers.NewGatewayEventMapper(mappers.WithLogger(logger))
	httpRouteEventMapper := mappers.NewHTTPRouteEventMapper(mappers.WithLogger(logger))

	return ctrl.NewControllerManagedBy(mgr).
		For(&kuadrantv1beta1.MyPolicy{}).
		Watches(
			&source.Kind{Type: &gatewayapiv1beta1.HTTPRoute{}},
			handler.EnqueueRequestsFromMapFunc(httpRouteEventMapper.MapToRateLimitPolicy),
		).
		Watches(
			&source.Kind{Type: &gatewayapiv1beta1.Gateway{}},
			handler.EnqueueRequestsFromMapFunc(gatewayEventMapper.MapToRateLimitPolicy),
		).
		Complete(r)
}
```

### Controller-runtime client extension functions

**`NamespacedNameToObjectKey`**<br/>
Converts a `namespace/name` string to a Kubernetes object key.

**`ReadAnnotationsFromObject`**<br/>
Reads the annotations of a Kubernetes object and returns them as a map. If the object has no annotations, it returns an empty map.

This is similar to [`ObjectMeta.GetAnnotations()`](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta.GetAnnotations), except that it allows safely using the returned object without risk of null-pointer exceptions.

### Generic Golang util functions

Basic Golang helper functions leveraging generics to handle slices.

- **`Contains[T comparable](slice []T, target T) bool`**
- **`Find[T any](slice []T, match func(T) bool) (*T, bool)`**
- **`IndexOf[T comparable](slice []T, target T) int`**
- **`Map[T, U any](slice []T, f func(T) U) []U`**
- **`SliceCopy[T any](s1 []T) []T`**
- **`ReverseSlice[T any](input []T) []T`**
