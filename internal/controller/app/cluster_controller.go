/*
Copyright 2024.

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

package app

import (
	"context"
	"errors"
	appv1 "example.org/multi-clusters/api/app/v1"
	"example.org/multi-clusters/common"
	"example.org/multi-clusters/pkg/provider"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// ClusterReconciler reconciles a Cluster object
type ClusterReconciler struct {
	client.Client
	Scheme  *runtime.Scheme
	Eventer record.EventRecorder
	Log     logr.Logger
}

//+kubebuilder:rbac:groups=app.example.org,resources=clusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.example.org,resources=clusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.example.org,resources=clusters/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=events,verbs=get;list;watch;create;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *ClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithName(req.NamespacedName.String())
	logger.Info("handling change...")
	cluster := &appv1.Cluster{}
	if err := r.Get(ctx, req.NamespacedName, cluster); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !cluster.ObjectMeta.DeletionTimestamp.IsZero() {
		logger.Info("cleaning resource deleting cluster")
		controllerutil.RemoveFinalizer(cluster, "operator-controller")
		err := r.Update(context.TODO(), cluster)
		if err != nil {
			logger.Error(err, "add finalizer fail")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	logger.V(0).Info("cluster info",
		"version", cluster.Spec.Version,
		"CNI", cluster.Spec.CNI,
		"master", cluster.Spec.MasterSize,
		"worker", cluster.Spec.WorkerSize)

	if cluster.ObjectMeta.Finalizers == nil {
		controllerutil.AddFinalizer(cluster, "operator-controller")
		err := r.Update(context.TODO(), cluster)
		if err != nil {
			logger.Error(err, "add finalizer fail")
		}
	}
	clusterProvider := provider.GetClusterProvider(cluster.Spec.Provider)
	c, err := clusterProvider.GetCluster(cluster.Name)
	if err != nil {
		if !errors.Is(err, common.ClusterNotFoundError) {
			logger.Error(err, "Get cluster error")
			return ctrl.Result{}, err
		}
		// not found, need to create one
		err = clusterProvider.CreateCluster(cluster)
		if err != nil {
			// 清理残留资源？更新condition？
			logger.Error(err, "Create cluster error")
			return ctrl.Result{}, err
		}
	}
	// 获取c中的集群健康、节点等信息，和当前cluster比对，不一致时更新
	// 节点ip不一致时更新kind集群，集群健康不一致时更新cr资源
	inspect(cluster, c)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.Cluster{}).
		Complete(r)
}

func inspect(cluster *appv1.Cluster, c interface{}) {

}
