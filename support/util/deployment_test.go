package util

import (
	"context"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIsDeploymentReady(t *testing.T) {
	client := fake.NewClientBuilder().WithRuntimeObjects(
		&appsv1.Deployment{
			ObjectMeta: v1.ObjectMeta{
				Name: "positive-path",
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: pointer.Int32Ptr(3),
			},
			Status: appsv1.DeploymentStatus{
				UpdatedReplicas:   3,
				AvailableReplicas: 3,
			},
		},
		&appsv1.Deployment{
			ObjectMeta: v1.ObjectMeta{
				Name: "negative-path-1",
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: pointer.Int32Ptr(3),
			},
			Status: appsv1.DeploymentStatus{
				UpdatedReplicas:   2,
				AvailableReplicas: 3,
			},
		},
		&appsv1.Deployment{
			ObjectMeta: v1.ObjectMeta{
				Name: "negative-path-2",
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: pointer.Int32Ptr(3),
			},
			Status: appsv1.DeploymentStatus{
				UpdatedReplicas:   3,
				AvailableReplicas: 2,
			},
		},
	).Build()

	// Positive path - all replicas updated and ready
	result, err := IsDeploymentReady(context.TODO(), client, &appsv1.Deployment{ObjectMeta: v1.ObjectMeta{Name: "positive-path"}})
	if !result || err != nil {
		t.Errorf("expected result %t, got %t: %v", true, result, err)
	}

	// Negative path - replicas not yet updated
	result, err = IsDeploymentReady(context.TODO(), client, &appsv1.Deployment{ObjectMeta: v1.ObjectMeta{Name: "negative-path-1"}})
	if result || err == nil {
		t.Errorf("expected result %t, got %t", false, result)
	}

	// Negative path - replicas not yet available
	result, err = IsDeploymentReady(context.TODO(), client, &appsv1.Deployment{ObjectMeta: v1.ObjectMeta{Name: "negative-path-2"}})
	if result || err == nil {
		t.Errorf("expected result %t, got %t", false, result)
	}
}
