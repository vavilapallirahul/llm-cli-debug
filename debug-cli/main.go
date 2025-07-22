package debugcli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Run_cmd(clientset *kubernetes.Clientset) {
	var pod = &cobra.Command{
		Use:   "pod",
		Short: "name of pod comes here",
		Long:  `pto get pods name for debug purpose using logs`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ns, _ := cmd.Flags().GetString("namespace")
			pods, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}

			for _, val := range pods.Items {
				fmt.Println("pods name : ", val.Name)
				// fmt.Println("total spec : ", val)
			}
		},
	}

	pod.PersistentFlags().StringP("namespace", "n", "default", "namespace name in which pod is present")
	var rootCmd = &cobra.Command{Use: "debug"}
	rootCmd.AddCommand(pod)
	pod.Execute()
}
