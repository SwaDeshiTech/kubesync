package kubernetes

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func SyncSecret(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace, secretName string) error {
	ctx := context.Background()

	// Get the Secret from the source namespace
	secret, err := clientset.CoreV1().Secrets(sourceNamespace).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get Secret %s from namespace %s: %w", secretName, sourceNamespace, err)
	}

	// Create a new Secret in the target namespace with the same data
	targetSecret := secret.DeepCopy()
	targetSecret.Namespace = targetNamespace
	targetSecret.ObjectMeta.ResourceVersion = ""

	_, err = clientset.CoreV1().Secrets(targetNamespace).Create(ctx, targetSecret, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			// If the Secret already exists, update it
			_, err = clientset.CoreV1().Secrets(targetNamespace).Update(ctx, targetSecret, metav1.UpdateOptions{})
			if err != nil {
				return fmt.Errorf("failed to update Secret %s in namespace %s: %w", secretName, targetNamespace, err)
			}
			fmt.Printf("Secret %s synced from %s to %s\n", secretName, sourceNamespace, targetNamespace)
		} else {
			return fmt.Errorf("failed to create Secret %s in namespace %s: %w", secretName, targetNamespace, err)
		}
	} else {
		fmt.Printf("Secret %s synced from %s to %s\n", secretName, sourceNamespace, targetNamespace)
	}

	return nil
}
