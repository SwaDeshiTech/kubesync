package kubernetes

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func SyncConfigMap(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace, configMapName string) error {
	ctx := context.Background()
	var gracePeriodSeconds int64
	gracePeriodSeconds = 0

	// Get the ConfigMap from the source namespace
	configMap, err := clientset.CoreV1().ConfigMaps(sourceNamespace).Get(ctx, configMapName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get ConfigMap %s from namespace %s: %w", configMapName, sourceNamespace, err)
	}

	// Create a new ConfigMap in the target namespace with the same data
	targetConfigMap := configMap.DeepCopy()
	targetConfigMap.Namespace = targetNamespace

	err = clientset.CoreV1().ConfigMaps(targetNamespace).Delete(ctx, configMapName, metav1.DeleteOptions{GracePeriodSeconds: &gracePeriodSeconds})
	if err != nil {
		return fmt.Errorf("failed to delete ConfigMap %s in namespace %s: %w", configMapName, targetNamespace, err)
	}

	_, err = clientset.CoreV1().ConfigMaps(targetNamespace).Create(ctx, targetConfigMap, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to sync ConfigMap %s in namespace %s: %w", configMapName, targetNamespace, err)
	}

	fmt.Printf("ConfigMap %s synced from %s to %s\n", configMapName, sourceNamespace, targetNamespace)

	return nil
}
