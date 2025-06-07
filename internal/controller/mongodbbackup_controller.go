/*
Copyright 2025 rickyalex88@gmail.com.

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

package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/rickyalex/mongodb-backup-operator/api/v1"
)

// MongoDBBackupReconciler reconciles a MongoDBBackup object
type MongoDBBackupReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=backup.rilex-workspace.com,resources=mongodbbackups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=backup.rilex-workspace.com,resources=mongodbbackups/status,verbs=get;update;patch

func (r *MongoDBBackupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("mongodbbackup", req.NamespacedName)

	var backup v1.MongoDBBackup
	if err := r.Get(ctx, req.NamespacedName, &backup); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	log.Info("Starting backup", "uri", backup.Spec.DatabaseURI)

	// Simulate backup logic
	backupPath := fmt.Sprintf("%s/%s", backup.Spec.BackupPath, backup.Name)
	log.Info("Simulating mongodump", "output", backupPath)

	// Update status
	backup.Status.LastBackupTime = metav1.Time{Time: time.Now()}
	if err := r.Status().Update(ctx, &backup); err != nil {
		log.Error(err, "unable to update status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{RequeueAfter: 24 * time.Hour}, nil
}

func (r *MongoDBBackupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.MongoDBBackup{}).
		Complete(r)
}

