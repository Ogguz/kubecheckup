package k8s

import (
	"fmt"
	"github.com/Ogguz/kubecheckup/model"
	log "github.com/sirupsen/logrus"
	"github.com/tcnksm/go-latest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"

	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func RunAllTheTests(c *model.Config)  {
	k := initApiConnection(c)
	var result string // TODO send notification if return is false
	// TODO add go routine
	result,_ = checkKubernetesVersion(k)
	fmt.Println(result)

}

// initApiConnection reads kubeconfig file and returns clientset for api connection.
func initApiConnection(c *model.Config) *kubernetes.Clientset {

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

// checkKubernetesVersion gets the latest stable kubernetes version from github.com/kubernetes/kubernetes and compares
// it with the current one.
func checkKubernetesVersion(k *kubernetes.Clientset) (string, bool) {
	githubTag := &latest.GithubTag{
		Owner: "kubernetes",
		Repository: "kubernetes",
		TagFilterFunc: func(githubTag string) bool {
			if strings.Contains(githubTag,"alpha") {
				return false
			} else if strings.Contains(githubTag,"rc") {
				return false
			} else if strings.Contains(githubTag,"beta") {
				return false
			}
			return true
		},
	}

	latestVersion, err := k.ServerVersion()
	if err != nil {
		log.Error("Not able to get kubernetes cluster version", err)
	}

	res, _ := latest.Check(githubTag, fmt.Sprint(latestVersion))
	if res.Outdated {
		output :=  fmt.Sprint(latestVersion) + "is not latest, you should upgrade to " + res.Current
		return output, false
	}
	return "Kubernetes is up to date.", true
}
