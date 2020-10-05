package configurationresource

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateSecretSpec ...
func GenerateSecretSpec() {
	secret := corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "secret-data",
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"username": []byte("dGVzdA=="), // "test" in base64
			"password": []byte("MTIzNA=="), // "1234" in base64
		},
	}

	fmt.Printf("%+v", secret)
	// return configMap
}
