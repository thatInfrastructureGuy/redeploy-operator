package stub

import (
	"context"
	"fmt"
	"time"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/thatinfraguy/redeploy-operator/pkg/apis/app/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Redeploy:
		if o.Spec.RedeployNeeded && len(o.Spec.DeploymentName) > 0 && len(o.Spec.DeploymentNamespace) > 0 {
			fmt.Println("Redeploy is set to TRUE!!!")
			currentTime := time.Now()
			redeployDate := currentTime.Format("2006-01-02-15_04_05")

			o.Spec.RedeployNeeded = false

			o.Status.Status = "REDEPLOYED"
			o.Status.Date = redeployDate

			sdk.Update(o)

			config, err := rest.InClusterConfig()
			if err != nil {
				panic(err.Error())
			}
			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}

			deploymentsClient := clientset.AppsV1().Deployments(o.Spec.DeploymentNamespace)
			fmt.Printf("Redeploying %s at %s", o.Spec.DeploymentName, redeployDate)

			retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				// Retrieve the latest version of Deployment before attempting update
				// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
				result, getErr := deploymentsClient.Get(o.Spec.DeploymentName, metav1.GetOptions{})
				if getErr != nil {
					fmt.Printf("Failed to get latest version of Deployment: %v", getErr)
					return getErr
				}

				result.Spec.Template.ObjectMeta.Labels["redeployed"] = redeployDate
				_, updateErr := deploymentsClient.Update(result)
				return updateErr

			})

			if retryErr != nil {
				fmt.Printf("Update failed: %v", retryErr)
			} else {
				fmt.Printf("Redeployed %s", o.Spec.DeploymentName)
			}

		} else {
			fmt.Println("Nothing to Redeploy!")
		}
	}
	return nil
}
