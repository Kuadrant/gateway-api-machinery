package reconcilers

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/log"
	gatewayapiv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func TestReconcileTargetBackReference(t *testing.T) {
	var (
		namespace             = "operator-unittest"
		routeName             = "my-route"
		annotationName string = "some-annotation"
	)
	baseCtx := context.Background()
	ctx := logr.NewContext(baseCtx, log.Log)

	s := scheme.Scheme
	err := appsv1.AddToScheme(s)
	if err != nil {
		t.Fatal(err)
	}
	err = gatewayapiv1beta1.AddToScheme(s)
	if err != nil {
		t.Fatal(err)
	}

	policyKey := client.ObjectKey{Name: "someName", Namespace: "someNamespace"}

	existingRoute := &gatewayapiv1beta1.HTTPRoute{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "gateway.networking.k8s.io/v1beta1",
			Kind:       "HTTPRoute",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      routeName,
			Namespace: namespace,
		},
		Spec: gatewayapiv1beta1.HTTPRouteSpec{
			CommonRouteSpec: gatewayapiv1beta1.CommonRouteSpec{
				ParentRefs: []gatewayapiv1beta1.ParentReference{
					{
						Name: "gwName",
					},
				},
			},
		},
		Status: gatewayapiv1beta1.HTTPRouteStatus{
			RouteStatus: gatewayapiv1beta1.RouteStatus{
				Parents: []gatewayapiv1beta1.RouteParentStatus{
					{
						ParentRef: gatewayapiv1beta1.ParentReference{
							Name: "gwName",
						},
						Conditions: []metav1.Condition{
							{
								Type:   "Accepted",
								Status: metav1.ConditionTrue,
							},
						},
					},
				},
			},
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{existingRoute}

	// Create a fake client to mock API calls.
	cl := fake.NewFakeClient(objs...)

	targetRefReconciler := TargetRefReconciler{
		Client: cl,
	}

	err = targetRefReconciler.ReconcileTargetBackReference(ctx, policyKey, existingRoute, annotationName)
	if err != nil {
		t.Fatal(err)
	}

	res := &gatewayapiv1beta1.HTTPRoute{}
	err = cl.Get(ctx, client.ObjectKey{Name: routeName, Namespace: namespace}, res)
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatal("res is nil")
	}

	if len(res.GetAnnotations()) == 0 {
		t.Fatal("annotations are empty")
	}

	val, ok := res.GetAnnotations()[annotationName]
	if !ok {
		t.Fatal("expected annotation not found")
	}

	if val != policyKey.String() {
		t.Fatalf("annotation value (%s) does not match expected (%s)", val, policyKey.String())
	}
}

func TestDeleteTargetBackReference(t *testing.T) {
	var (
		namespace             = "operator-unittest"
		routeName             = "my-route"
		annotationName string = "some-annotation"
	)
	baseCtx := context.Background()
	ctx := logr.NewContext(baseCtx, log.Log)

	s := scheme.Scheme
	err := appsv1.AddToScheme(s)
	if err != nil {
		t.Fatal(err)
	}
	err = gatewayapiv1beta1.AddToScheme(s)
	if err != nil {
		t.Fatal(err)
	}

	policyKey := client.ObjectKey{Name: "someName", Namespace: "someNamespace"}

	existingRoute := &gatewayapiv1beta1.HTTPRoute{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "gateway.networking.k8s.io/v1beta1",
			Kind:       "HTTPRoute",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      routeName,
			Namespace: namespace,
			Annotations: map[string]string{
				annotationName: "annotationValue",
			},
		},
		Spec: gatewayapiv1beta1.HTTPRouteSpec{
			CommonRouteSpec: gatewayapiv1beta1.CommonRouteSpec{
				ParentRefs: []gatewayapiv1beta1.ParentReference{
					{
						Name: "gwName",
					},
				},
			},
		},
		Status: gatewayapiv1beta1.HTTPRouteStatus{
			RouteStatus: gatewayapiv1beta1.RouteStatus{
				Parents: []gatewayapiv1beta1.RouteParentStatus{
					{
						ParentRef: gatewayapiv1beta1.ParentReference{
							Name: "gwName",
						},
						Conditions: []metav1.Condition{
							{
								Type:   "Accepted",
								Status: metav1.ConditionTrue,
							},
						},
					},
				},
			},
		},
	}

	// Objects to track in the fake client.
	objs := []runtime.Object{existingRoute}

	// Create a fake client to mock API calls.
	cl := fake.NewFakeClient(objs...)
	targetRefReconciler := TargetRefReconciler{
		Client: cl,
	}

	err = targetRefReconciler.DeleteTargetBackReference(ctx, policyKey, existingRoute, annotationName)
	if err != nil {
		t.Fatal(err)
	}

	res := &gatewayapiv1beta1.HTTPRoute{}
	err = cl.Get(ctx, client.ObjectKey{Name: routeName, Namespace: namespace}, res)
	if err != nil {
		t.Fatal(err)
	}

	if res == nil {
		t.Fatal("res is nil")
	}

	if len(res.GetAnnotations()) > 0 {
		_, ok := res.GetAnnotations()[annotationName]
		if ok {
			t.Fatal("expected annotation found and it should have been deleted")
		}
	}
}
