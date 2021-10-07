/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package postgresqldatabase

import (
	"time"

	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	tjcontroller "github.com/crossplane-contrib/terrajet/pkg/controller"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-azure/apis/postgresql/v1alpha1"
)

// Setup adds a controller that reconciles PostgresqlDatabase managed resources.
func Setup(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, s terraform.SetupFn, ws *terraform.WorkspaceStore, concurrency int) error {
	name := managed.ControllerName(v1alpha1.PostgresqlDatabaseGroupVersionKind.String())
	r := managed.NewReconciler(mgr,
		xpresource.ManagedKind(v1alpha1.PostgresqlDatabaseGroupVersionKind),
		managed.WithExternalConnecter(tjcontroller.NewConnector(mgr.GetClient(), ws, s)),
		managed.WithLogger(l.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(terraform.NewWorkspaceFinalizer(ws, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3*time.Minute),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{RateLimiter: rl, MaxConcurrentReconciles: concurrency}).
		For(&v1alpha1.PostgresqlDatabase{}).
		Complete(r)
}
