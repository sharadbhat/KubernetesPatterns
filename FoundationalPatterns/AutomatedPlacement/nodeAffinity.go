package automatedplacement

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateNodeAffinitySpec ...
func GenerateNodeAffinitySpec() {
	pod := corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "node-affinity",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "affinity-container",
					Image: "busybox",
					Args: []string{
						"/bin/sh",
						"-c",
						"echo \"Hello\"",
					},
				},
			},
			Affinity: &corev1.Affinity{
				NodeAffinity: &corev1.NodeAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
						NodeSelectorTerms: []corev1.NodeSelectorTerm{
							{
								MatchExpressions: []corev1.NodeSelectorRequirement{
									{
										Key:      "numberCores",
										Operator: corev1.NodeSelectorOpGt,
										Values:   []string{"1"},
									},
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
