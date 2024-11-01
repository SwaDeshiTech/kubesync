package kubernetes

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func SyncConfigMap(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace, configMapName string) error {
	ctx := context.Background()

	// Get the ConfigMap from the source namespace
	sourceConfigMap, err := clientset.CoreV1().ConfigMaps(sourceNamespace).Get(ctx, configMapName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get ConfigMap %s from namespace %s: %w", configMapName, sourceNamespace, err)
	}

	// Create a new ConfigMap in the target namespace with the same data
	targetConfigMap := sourceConfigMap.DeepCopy()
	targetConfigMap.Namespace = targetNamespace

	if !checkIfConfigMapExists(ctx, clientset, configMapName, targetNamespace) {
		targetConfigMap.ObjectMeta.ResourceVersion = ""
	}

	_, err = clientset.CoreV1().ConfigMaps(targetNamespace).Create(ctx, targetConfigMap, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			// If the ConfigMap already exists, update it
			_, err = clientset.CoreV1().ConfigMaps(targetNamespace).Update(ctx, targetConfigMap, metav1.UpdateOptions{})
			if err != nil {
				return fmt.Errorf("failed to update ConfigMap %s in namespace %s: %w", configMapName, targetNamespace, err)
			}
			fmt.Printf("ConfigMap %s synced from %s to %s\n", configMapName, sourceNamespace, targetNamespace)
		} else {
			return fmt.Errorf("failed to create ConfigMap %s in namespace %s: %w", configMapName, targetNamespace, err)
		}
	} else {
		fmt.Printf("ConfigMap %s synced from %s to %s\n", configMapName, sourceNamespace, targetNamespace)
	}

	return nil
}

func checkIfConfigMapExists(ctx context.Context, clientset *kubernetes.Clientset, configMapName, namespace string) bool {
	// Get the ConfigMap from the namespace
	_, err := clientset.CoreV1().ConfigMaps(namespace).Get(ctx, configMapName, metav1.GetOptions{})
	if err != nil && errors.IsAlreadyExists(err) {
		return true
	}
	return false
}
