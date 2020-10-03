package daemonservice

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateDaemonServiceSpec ...
func GenerateDaemonServiceSpec() {
	daemonService := appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "daemon-service",
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"name": "daemon-label",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": "daemon-label",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "daemon-container",
							Image: "busybox",
							Args: []string{
								"/bin/sh",
								"-c",
								`
								while true;
								do
									echo "Running as a Daemon service";
									sleep 1;
								done;
								`,
							},
						},
					},
				},
			},
		},
	}

	fmt.Printf("%+v", daemonService)
	// return daemonService
}
