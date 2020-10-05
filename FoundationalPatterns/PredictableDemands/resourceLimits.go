package predictabledemands

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateResourceLimitSpec ...
func GenerateResourceLimitSpec() {
	pod := corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "resource-limit-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "limit-container",
					Image: "busybox",
					Args: []string{
						"/bin/sh",
						"-c",
						`while true;
							do
								echo "Running forever";
								sleep 1;
							done;
						`,
					},
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    *resource.NewMilliQuantity(100, resource.DecimalSI),
							corev1.ResourceMemory: *resource.NewMilliQuantity(100, resource.DecimalSI),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    *resource.NewMilliQuantity(200, resource.DecimalSI),
							corev1.ResourceMemory: *resource.NewMilliQuantity(200, resource.DecimalSI),
						},
					},
				},
			},
		},
	}

	fmt.Printf("%+v", pod)
	// return pod
}
