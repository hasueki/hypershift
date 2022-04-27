package util

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func IsDeploymentReady(ctx context.Context, c crclient.Client, deployment *appsv1.Deployment) (bool, error) {
	if err := c.Get(ctx, crclient.ObjectKeyFromObject(deployment), deployment); err != nil {
		return false, fmt.Errorf("failed to fetch %s deployment: %w", deployment.Name, err)
	}

	expectedReplicas := deployment.Spec.Replicas
	if deployment.Status.UpdatedReplicas != *expectedReplicas {
		return false, fmt.Errorf("replicas not yet updated for %s deployment", deployment.Name)
	}
	if deployment.Status.AvailableReplicas != *expectedReplicas {
		return false, fmt.Errorf("replicas not yet available for %s deployment", deployment.Name)
	}

	return true, nil
}
