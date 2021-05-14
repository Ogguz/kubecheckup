package k8s

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func checkCoreComponenets(k kubernetes.Clientset) (string, bool) {

	coreComponents, err := k.CoreV1().ComponentStatuses().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		log.Fatal("Could not get core components statuses ", err)
		return "Couldn't get core components statuses...",false
	}

	listOfUnhealthyComponents := make([]string,0)

	for _,i := range coreComponents.Items {
		for _,s := range i.Conditions {
			if s.Status != "true" {
				listOfUnhealthyComponents = append(listOfUnhealthyComponents,i.GetName())
			}
		}
	}

	if len(listOfUnhealthyComponents)>=1 {
		for _, i := range listOfUnhealthyComponents {
			fmt.Println(i)
		}
		return "Unhealthy core components!!",false
	}

	return "OK, core components are healthy", true
}
