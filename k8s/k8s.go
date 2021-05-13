package k8s

import (
	"github.com/Ogguz/kubecheckup/model"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

// InitApiConnection reads kubeconfig file and returns clientset for api connection.
func InitApiConnection(c *model.Config) *kubernetes.Clientset {

	kubeconfig := c.Kubernetes.ConfigFile

	log.Debugf("Reading kubeconfig file from %s",kubeconfig)
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	log.Debugf("Successfully read %s file, creating the clientset...",kubeconfig)
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	log.Debug("Clientset creation succeed")

	return clientset
}
