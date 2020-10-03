package envvarconfiguration

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// GenerateEnvVarConfigurationSpec ...
func GenerateEnvVarConfigurationSpec() {
	pod := &corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  "main-container",
				Image: "busybox",
				Env: []corev1.EnvVar{
					{
						Name:  "VALUE_1",
						Value: "This is stored in VALUE_1",
					},
				},
				Args: []string{
					"/bin/sh",
					"-c",
					"Value 1 -> $VALUE_1",
				},
			},
		},
		RestartPolicy: "Never",
	}

	fmt.Printf("%+v", pod)
	// return pod
}
