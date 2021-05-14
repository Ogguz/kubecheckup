package k8s

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func checkDeployments(k *kubernetes.Clientset) (string, bool) {

	deployments, err := k.AppsV1().Deployments("").List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		log.Fatal("Could not get deployment list ", err)
	}

	listOfOrphanDeployments := make([]string,0)
	listOfLeftoverDeployments := make([]string,0)

	for _,i := range deployments.Items {
		if i.Status.Replicas == 0 && i.Status.AvailableReplicas == 0 {
			listOfLeftoverDeployments = append(listOfLeftoverDeployments,i.GetName())
		} else if i.Status.Replicas > 0 && i.Status.AvailableReplicas == 0 {
			listOfOrphanDeployments = append(listOfOrphanDeployments,i.GetName())
		}
	}

    print := func(s []string) {
    	if len(s) >= 1 {
    		for _,i := range s {
    			fmt.Println(i)
			}
		}
	}

	print(listOfLeftoverDeployments)
    print(listOfOrphanDeployments)

    if len(listOfOrphanDeployments) == 0 && len(listOfLeftoverDeployments) == 0 {
		return "OK, No leftover or orphan deployment",true
	}

	return "Found leftover/orphan deployments", false
}
