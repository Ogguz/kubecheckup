package k8s

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// checkReplicasets Checks for orphan replicasets
func checkReplicasets(k *kubernetes.Clientset) (string, bool) {

	replicasets, err := k.AppsV1().ReplicaSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Could not get replicaset list ", err)
	}

	listOfOrphanReplicasets := make([]string, 0)
	listOfLeftoverReplicasets := make([]string, 0)

	for _, i := range replicasets.Items {
		if i.Status.Replicas == 0 && i.Status.AvailableReplicas == 0 {
			listOfLeftoverReplicasets = append(listOfLeftoverReplicasets, i.GetName())
		} else if i.Status.Replicas > 0 && i.Status.AvailableReplicas == 0 {
			listOfOrphanReplicasets = append(listOfOrphanReplicasets, i.GetName())
		}
	}

	print := func(s []string) {
		if len(s) >= 1 {
			for _, i := range s {
				fmt.Println(i)
			}
		}
	}

	print(listOfLeftoverReplicasets)
	print(listOfOrphanReplicasets)

	if len(listOfOrphanReplicasets) == 0 && len(listOfLeftoverReplicasets) == 0 {
		return "OK, No leftover or orphan replicaset", true
	}

	return "Found leftover/orphan replicasets", false
}
