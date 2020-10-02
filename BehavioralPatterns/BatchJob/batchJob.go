package batchjob

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateBatchJobSpec ...
func GenerateBatchJobSpec() {
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "batch-job",
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "batch-container",
							Image: "busybox",
							Args: []string{
								"/bin/sh",
								"-c",
								"echo \"Executed at $(date)\"",
							},
						},
					},
				},
			},
		},
	}

	fmt.Printf("%+v", job)
	// return job
}
