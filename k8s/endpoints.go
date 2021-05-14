package k8s

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// orphanEndpoints gets a client for k8s and scans through all endpoints to check if they are leftover.
func orphanEndpoints(k *kubernetes.Clientset) (string,bool){

	// TODO user should be able to choose the namespace via config file
	endpoints, err := k.CoreV1().Endpoints("").List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		log.Fatal("Could not get endpoint list ", err)
		return "Couldn't get endpoint list...",false
	}

	listOfOrphanEndpoints := make([]string,0)

	for _,i := range endpoints.Items {
		if len(i.Subsets) == 0 {
			listOfOrphanEndpoints = append(listOfOrphanEndpoints,i.GetName())
		}
	}

	if len(listOfOrphanEndpoints)>=1 {
		for _,i := range listOfOrphanEndpoints {
			fmt.Println(i)  // TODO print the namespace as well
		}
		return "Found orphan endpoints...", false
	}
	return "OK, no orphan endpoints",true
}
