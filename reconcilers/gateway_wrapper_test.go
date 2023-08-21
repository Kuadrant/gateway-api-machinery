package reconcilers

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gatewayapiv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	"github.com/kuadrant/gateway-api-machinery/common"
)

func TestGatewayWrapperKey(t *testing.T) {
	gw := GatewayWrapper{
		Gateway: &gatewayapiv1beta1.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   "gw-ns",
				Name:        "gw-1",
				Annotations: map[string]string{"kuadrant.io/testpolicies": `[{"Namespace":"app-ns","Name":"policy-1"},{"Namespace":"app-ns","Name":"policy-2"}]`},
			},
		},
		Referrer: &common.PolicyKindStub{},
	}
	if gw.Key().Namespace != "gw-ns" || gw.Key().Name != "gw-1" {
		t.Fail()
	}
}

func TestGatewayWrapperContainsPolicy(t *testing.T) {
	gw := GatewayWrapper{
		Gateway: &gatewayapiv1beta1.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   "gw-ns",
				Name:        "gw-1",
				Annotations: map[string]string{"kuadrant.io/testpolicies": `[{"Namespace":"app-ns","Name":"policy-1"},{"Namespace":"app-ns","Name":"policy-2"}]`},
			},
		},
		Referrer: &common.PolicyKindStub{},
	}
	if !gw.ContainsPolicy(client.ObjectKey{Namespace: "app-ns", Name: "policy-1"}) {
		t.Error("GatewayWrapper.ContainsPolicy() should contain app-ns/policy-1")
	}
	if !gw.ContainsPolicy(client.ObjectKey{Namespace: "app-ns", Name: "policy-2"}) {
		t.Error("GatewayWrapper.ContainsPolicy() should contain app-ns/policy-1")
	}
	if gw.ContainsPolicy(client.ObjectKey{Namespace: "app-ns", Name: "policy-3"}) {
		t.Error("GatewayWrapper.ContainsPolicy() should not contain app-ns/policy-1")
	}
}

func TestGatewayWrapperAddPolicy(t *testing.T) {
	gateway := gatewayapiv1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   "gw-ns",
			Name:        "gw-1",
			Annotations: map[string]string{"kuadrant.io/testpolicies": `[{"Namespace":"app-ns","Name":"policy-1"},{"Namespace":"app-ns","Name":"policy-2"}]`},
		},
	}
	gw := GatewayWrapper{
		Gateway:  &gateway,
		Referrer: &common.PolicyKindStub{},
	}
	if gw.AddPolicy(client.ObjectKey{Namespace: "app-ns", Name: "policy-1"}) {
		t.Error("GatewayWrapper.AddPolicy() expected to return false")
	}
	if !gw.AddPolicy(client.ObjectKey{Namespace: "app-ns", Name: "policy-3"}) {
		t.Error("GatewayWrapper.AddPolicy() expected to return true")
	}
	if gw.Annotations["kuadrant.io/testpolicies"] != `[{"Namespace":"app-ns","Name":"policy-1"},{"Namespace":"app-ns","Name":"policy-2"},{"Namespace":"app-ns","Name":"policy-3"}]` {
		t.Error("GatewayWrapper.AddPolicy() expected to have added policy ref to the annotations")
	}
}

func TestGatewayDeletePolicy(t *testing.T) {
	gateway := gatewayapiv1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   "gw-ns",
			Name:        "gw-1",
			Annotations: map[string]string{"kuadrant.io/testpolicies": `[{"Namespace":"app-ns","Name":"policy-1"},{"Namespace":"app-ns","Name":"policy-2"}]`},
		},
	}
	gw := GatewayWrapper{
		Gateway:  &gateway,
		Referrer: &common.PolicyKindStub{},
	}
	if !gw.DeletePolicy(client.ObjectKey{Namespace: "app-ns", Name: "policy-1"}) {
		t.Error("GatewayWrapper.DeletePolicy() expected to return true")
	}
	if gw.DeletePolicy(client.ObjectKey{Namespace: "app-ns", Name: "policy-3"}) {
		t.Error("GatewayWrapper.DeletePolicy() expected to return false")
	}
	if gw.Annotations["kuadrant.io/testpolicies"] != `[{"Namespace":"app-ns","Name":"policy-2"}]` {
		t.Error("GatewayWrapper.DeletePolicy() expected to have deleted policy ref from the annotations")
	}
}

func TestBackReferencesFromGatewayWrapper(t *testing.T) {
	gw := GatewayWrapper{
		Gateway: &gatewayapiv1beta1.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   "gw-ns",
				Name:        "gw-1",
				Annotations: map[string]string{"kuadrant.io/testpolicies": `[{"Namespace":"app-ns","Name":"policy-1"},{"Namespace":"app-ns","Name":"policy-2"}]`},
			},
		},
		Referrer: &common.PolicyKindStub{},
	}
	refs := common.Map(common.BackReferencesFromObject(gw.Gateway, gw.Referrer), func(ref client.ObjectKey) string { return ref.String() })
	if !common.Contains(refs, "app-ns/policy-1") {
		t.Error("GatewayWrapper.PolicyRefs() should contain app-ns/policy-1")
	}
	if !common.Contains(refs, "app-ns/policy-2") {
		t.Error("GatewayWrapper.PolicyRefs() should contain app-ns/policy-2")
	}
	if len(refs) != 2 {
		t.Fail()
	}
}
