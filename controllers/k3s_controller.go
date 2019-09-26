/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	provisionerv1 "github.com/kevin-zhaoshuai/k3s-operator/api/v1"
	"github.com/kevin-zhaoshuai/k3s-operator/provisioner"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

// K3sReconciler reconciles a K3s object
type K3sReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=provisioner.k3s.operator,resources=k3s,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=provisioner.k3s.operator,resources=k3s/status,verbs=get;update;patch

func (r *K3sReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("k3s", req.NamespacedName)

	// your logic here
	setupLog.Info("Just go for test")
	edgeNode := &provisionerv1.K3s{}
	err := r.Get(ctx, req.NamespacedName, edgeNode)
	if err != nil {
		setupLog.Error(err, "Fail to get edgeNode CRD")
	}
	// Initial the Edge Node status
	if edgeNode.Status.LastUpdateTimestamp == nil {
		edgeNode.Status.Type = edgeNode.Spec.Type
		edgeNode.Status.Phase = provisionerv1.ProvisionInit
		now := metav1.Now()
		edgeNode.Status.LastUpdateTimestamp = &now
		err = r.Update(ctx, edgeNode)
		return ctrl.Result{}, nil
	}

	if edgeNode.Status.Phase == provisionerv1.ProvisionSucceed {
		return ctrl.Result{}, nil
	}

	if edgeNode.Status.Phase == provisionerv1.ProvisionFailed {
		return ctrl.Result{}, nil
	}
	if edgeNode.Status.Phase == provisionerv1.ProvisionInit {
		edgeNode.Status.Type = edgeNode.Spec.Type
		errProvisioner := provisioner.ProvisionEdgeNode(*edgeNode)
		if errProvisioner != nil {
			setupLog.Error(errProvisioner, "Provision failed")
			edgeNode.Status.Phase = provisionerv1.ProvisionFailed
		} else {
			edgeNode.Status.Phase = provisionerv1.ProvisionSucceed
		}
		now := metav1.Now()
		edgeNode.Status.LastUpdateTimestamp = &now
		err = r.Update(ctx, edgeNode)
	}
	return ctrl.Result{}, nil
}

func (r *K3sReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&provisionerv1.K3s{}).
		Complete(r)
}
