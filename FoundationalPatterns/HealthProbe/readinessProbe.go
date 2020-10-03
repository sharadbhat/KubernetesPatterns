package healthprobe

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// GenerateReadinessProbeSpec ...
func GenerateReadinessProbeSpec() {
	pod := corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  "ready-container",
				Image: "busybox",
				Args: []string{
					"/bin/sh",
					"-c",
					`touch /tmp/healthy;
					sleep 30;
					rm -rf /tmp/healthy;
					sleep 600
					`,
				},
				ReadinessProbe: &corev1.Probe{
					Handler: corev1.Handler{
						Exec: &corev1.ExecAction{
							Command: []string{
								"cat",
								"/tmp/healthy",
							},
						},
					},
				},
			},
		},
	}

	fmt.Printf("%+v", pod)
	// return pod
}
