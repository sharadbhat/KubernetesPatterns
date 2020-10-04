package sidecar

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// GenerateSidecarSpec ...
func GenerateSidecarSpec() {
	pod := &corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  "main-container",
				Image: "busybox",
				Args: []string{
					"/bin/sh",
					"-c",
					`i=0;
					while true;
					do
					echo "Logging line: $i" >> /var/log/file.log;
					i=$((i+1));
					sleep 1;
					done`,
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "sharedlog",
						MountPath: "/var/log",
					},
				},
			},
			{
				Name:  "sidecar-container",
				Image: "busybox",
				Args: []string{
					"/bin/sh",
					"-c",
					"tail -f /var/log/file.log",
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "sharedlog",
						MountPath: "/var/log",
					},
				},
			},
		},
		Volumes: []corev1.Volume{
			{
				Name: "sharedlog",
			},
		},
	}

	fmt.Printf("%+v", pod)
	// return pod
}
