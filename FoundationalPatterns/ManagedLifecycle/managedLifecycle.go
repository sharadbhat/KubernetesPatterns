package managedlifecycle

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateManagedLifecycleSpec ...
func GenerateManagedLifecycleSpec() {
	pod := corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "lifecycle-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "lifecycle-container",
					Image: "busybox",
					Args: []string{
						"/bin/sh",
						"-c",
						`while true;
					do
						echo "Running forever";
						sleep 1;
					done`,
					},
					Lifecycle: &corev1.Lifecycle{
						PostStart: &corev1.Handler{
							Exec: &corev1.ExecAction{
								Command: []string{
									"/bin/sh",
									"-c",
									"touch /var/file.txt",
								},
							},
						},
						PreStop: &corev1.Handler{
							Exec: &corev1.ExecAction{
								Command: []string{
									"/bin/sh",
									"-c",
									"rm /var/file.txt",
								},
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
