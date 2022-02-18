package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func config() *kubernetes.Clientset {

	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}

	clientset := kubernetes.NewForConfigOrDie(config)

	return clientset

}
func nodes(clientset kubernetes.Clientset) {
	nodelist, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, n := range nodelist.Items {
		fmt.Println(n.Name)
	}
}

func GetSecret(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("prod").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=auth-prod",
	})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

	for _, pod := range pods.Items {
		data, _ := json.Marshal(pod.Name)
		fmt.Printf(string(data))

		err = clientset.CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})

		if err != nil {
			panic(err)

		}

	}
}
// print all pods in namespace
func listpod(clientset *kubernetes.Clientset) {
	podlist, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range podlist.Items {
		fmt.Println(pod.Name)

	}
}

//
///	podselector,  err:= metav1.ParseToLabelSelector("app=excel")
//if err != nil {
//	fmt.Printf(podselector.MatchLabels["app=excel"])
//}

//}

//func reload(clientset *kubernetes.Clientset) {

//	podlist, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
//	if err != nil {
//		panic(err)
//	}
///
//for _, pod := range podlist.Items {
//	fmt.Println(pod.Name)
//}

//err = clientset.CoreV1().Pods("dev").Delete(pod.Name, &metav1.DeleteOptions{})

//if err != nil {
//	panic(err)
//}

//}

func main() {
	var clientset = config()
	GetSecret(*&clientset)

}
