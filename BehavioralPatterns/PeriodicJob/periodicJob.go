package periodicjob

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GeneratePeriodicJobSpec ...
func GeneratePeriodicJobSpec() {

	successfulJobHistoryLimit := int32(3)
	failedJobHistoryLimit := int32(3)
	startingDeadlineSeconds := int64(60)
	concurrenyPolicy := batchv1beta1.ConcurrencyPolicy("Forbid")

	cronJob := &batchv1beta1.CronJob{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "CronJob",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "periodic-job",
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule:                   "*/1 * * * *",
			SuccessfulJobsHistoryLimit: &successfulJobHistoryLimit,
			FailedJobsHistoryLimit:     &failedJobHistoryLimit,
			StartingDeadlineSeconds:    &startingDeadlineSeconds,
			ConcurrencyPolicy:          concurrenyPolicy,
			JobTemplate: batchv1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "periodic-container",
									Image: "busybox",
									Args: []string{
										"/bin/sh",
										"-c",
										"echo \"Executed at $(date)\"",
									},
								},
							},
							RestartPolicy: "Never",
						},
					},
				},
			},
		},
	}

	fmt.Printf("%+v", cronJob)
	// return cronJob
}
