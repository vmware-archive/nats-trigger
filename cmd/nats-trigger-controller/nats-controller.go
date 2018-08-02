/*
Copyright (c) 2016-2017 Bitnami

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	kubelessutils "github.com/kubeless/kubeless/pkg/utils"
	"github.com/kubeless/nats-trigger/pkg/controller"
	natsutils "github.com/kubeless/nats-trigger/pkg/utils"
	"github.com/kubeless/nats-trigger/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nats-controller",
	Short: "NATS controller",
	Long: `NATS trigger CRD controller that watches for the creation/deletion/update events
				  of natstrigger API object from the Kubernetes API server and creates/deletes NATS subsciber to
				  requested topics. On recieving message from the topic from NATS fowards the message tp appropraiate
				  functions`,
	Run: func(cmd *cobra.Command, args []string) {

		kubelessClient, err := kubelessutils.GetFunctionClientInCluster()
		if err != nil {
			logrus.Fatalf("Cannot get kubeless CR API client: %v", err)
		}

		natsClient, err := natsutils.GetTriggerClientInCluster()
		if err != nil {
			logrus.Fatalf("Cannot get NATS trigger CR API client: %v", err)
		}

		natsTriggerCfg := controller.NatsTriggerConfig{
			KubeCli:        kubelessutils.GetClient(),
			TriggerClient:  natsClient,
			KubelessClient: kubelessClient,
		}

		natsTriggerController := controller.NewNatsTriggerController(natsTriggerCfg)

		stopCh := make(chan struct{})
		defer close(stopCh)

		go natsTriggerController.Run(stopCh)

		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)
		signal.Notify(sigterm, syscall.SIGINT)
		<-sigterm
	},
}

func main() {
	logrus.Infof("Running NATS controller version: %v", version.Version)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
