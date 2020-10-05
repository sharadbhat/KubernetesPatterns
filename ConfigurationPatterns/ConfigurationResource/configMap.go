package configurationresource

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateConfigMapSpec ...
func GenerateConfigMapSpec() {
	configMap := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "config-map-data",
		},
		Data: map[string]string{
			"sample-value": "from-config",
			"sample-data": `
			port=8080
			root=/
			`,
		},
	}

	fmt.Printf("%+v", configMap)
	// return configMap
}
