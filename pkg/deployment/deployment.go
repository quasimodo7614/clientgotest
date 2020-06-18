package deployment

import (
	"context"
	"flag"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	cliv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	"path/filepath"
)

var deployClient cliv1.DeploymentInterface

func init() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deployClient = clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
}

// ListDeployments returns all deployments.
func ListDeployments(ctx context.Context, count int) ([]appsv1.Deployment, error) {

	list, err := deployClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

// GetDeployment returns a deployment by id.
func GetDeployment(ctx context.Context, id string) (*appsv1.Deployment, error) {

	result, getErr := deployClient.Get("nginx-deployment", metav1.GetOptions{})
	if getErr != nil {
		return nil, getErr
	}
	return result, nil

}

// DeleteDeployment returns a deployment by id.
func DeleteDeployment(ctx context.Context, id string) (error) {
	err := deployClient.Delete("nginx-deployment", &metav1.DeleteOptions{})
	return err

}

// CreateDeployment create a deployment .
func CreateDeployment(ctx context.Context, a *appsv1.Deployment) (*appsv1.Deployment, error) {

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deployClient.Create(deployment)
	if err != nil {
		return nil, err
	}

	return result, nil

}

// UpdateDeployment returns a deployment by id.
func UpdateDeployment(ctx context.Context, name string) (*appsv1.Deployment, error) {
	fmt.Println(name)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deployClient.Get("nginx-deployment", metav1.GetOptions{})
		if getErr != nil {
			return getErr
		}

		result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
		_, updateErr := deployClient.Update(result)
		return updateErr
	})

	if retryErr != nil {
		return nil, retryErr
	}

	return nil, nil

}

func int32Ptr(i int32) *int32 { return &i }
