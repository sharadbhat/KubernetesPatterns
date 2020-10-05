package initcontainer

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateInitContainerSpec ...
func GenerateInitContainerSpec() {
	pod := corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "init-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "main-container",
					Image: "busybox",
					Args: []string{
						"/bin/sh",
						"-c",
						"ls /var/log",
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "sharedlog",
							MountPath: "/var/log",
						},
					},
				},
			},
			InitContainers: []corev1.Container{
				{
					Name:  "init-container",
					Image: "busybox",
					Args: []string{
						"/bin/sh",
						"-c",
						"touch /var/log/file.txt",
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
		},
	}

	fmt.Printf("%+v", pod)
	// return pod
}
