package k8s

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func checkNodeStatus(k *kubernetes.Clientset) (string, bool) {

	nodes, err := k.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		log.Fatal("Unable to list nodes ",err)
	}

	listUnreadyNodes := make([]string,0)

	for _,i := range nodes.Items {
		for _,s := range i.Status.Conditions {
			if s.Reason == "KubeletReady" && s.Status != "True" {
				listUnreadyNodes = append(listUnreadyNodes,i.GetName())
			}
		}
	}

	if len(listUnreadyNodes) >= 1 {
		for _,i := range listUnreadyNodes {
			fmt.Println(i)
		}
		return "Unready nodes detected!",false
	}

	return "OK, Nodes are healthy", true
}