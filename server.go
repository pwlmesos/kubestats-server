package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	// Kubernetes libraries
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {


	listen := flag.String("listen", ":8080", "Provide address:port to listen on")
	flag.Parse()

	log.Printf("Starting server on %s", *listen)
	http.HandleFunc("/", httpHandleRoot)
	http.HandleFunc("/ping", httpHandlePing)
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func getKubeConfig() string {

	    log.Printf("Enter 'getKubeConfig'")
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Printf("Error getting kubeconfig")
			return "<h1>Dude we ain't in a kubernetes cluster like we sposed to B\n<h1>" 
			//panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		returnString := ""
		//for {
			// get pods in all the namespaces by omitting namespace
			// Or specify namespace to get pods in particular namespace
			pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			returnString  = fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items))
	
			// Examples for error handling:
			// - Use helper functions e.g. errors.IsNotFound()
			// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
			_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
			if errors.IsNotFound(err) {
				fmt.Printf("Pod example-xxxxx not found in default namespace\n")
			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
				fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
			} else if err != nil {
				panic(err.Error())
			} else {
				fmt.Printf("Found example-xxxxx pod in default namespace\n")
			}
	
			//time.Sleep(10 * time.Second)
		//}
		return returnString
}


func httpHandleRoot(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/ping", http.StatusSeeOther)
}

func httpHandlePing(res http.ResponseWriter, req *http.Request) {
	log.Printf("/ping received")
	res.Header().Add("X-Powered-By", "golang")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, getKubeConfig())
}
