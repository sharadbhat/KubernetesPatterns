package initcontainer

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// GenerateInitContainerSpec ...
func GenerateInitContainerSpec() {
	pod := &corev1.PodSpec{
		Volumes: []corev1.Volume{
			{
				Name: "sharedlog",
			},
		},
		Containers: []corev1.Container{
			{
				Name:  "main-container",
				Image: "busybox",
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "sharedlog",
						MountPath: "/var/log",
					},
				},
				Args: []string{
					"/bin/sh",
					"-c",
					"touch /var/log/file.txt",
				},
			},
			{
				Name:  "sidecar-container",
				Image: "busybox",
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "sharedlog",
						MountPath: "/var/log",
					},
				},
				Args: []string{
					"/bin/sh",
					"-c",
					"ls /var/log",
				},
			},
		},
	}

	fmt.Printf("%+v", pod)
	// return pod
}
