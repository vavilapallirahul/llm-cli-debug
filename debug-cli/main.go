package debugcli

import (
	"bytes"
	"context"
	"encoding/json"
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
			po_name := args[0]

			fmt.Println(po_name)

			ns, _ := cmd.Flags().GetString("namespace")
			_, err := clientset.CoreV1().Namespaces().Get(context.Background(), ns, metav1.GetOptions{})

			if err != nil {
				fmt.Printf("no namespcae with name %s , is there \n", ns)
				panic(err.Error())
			}

			pod_info, err := clientset.CoreV1().Pods(ns).Get(context.Background(), po_name, metav1.GetOptions{})
			if err != nil {
				panic(err.Error())
			}

			pod_status_byte, err := json.Marshal(pod_info.Status)

			if err != nil {
				err.Error()
			}

			var prettyJSON bytes.Buffer
			json.Indent(&prettyJSON, pod_status_byte, "", "  ")

			fmt.Println(prettyJSON.String())

		},
	}

	pod.PersistentFlags().StringP("namespace", "n", "default", "namespace name in which pod is present")
	var rootCmd = &cobra.Command{Use: "debug"}
	rootCmd.AddCommand(pod)
	pod.Execute()
}
