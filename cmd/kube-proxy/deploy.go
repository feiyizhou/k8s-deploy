package kube_proxy

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
)

type kubeProxyService struct{}

func newKubeProxyService() *kubeProxyService {
	return &kubeProxyService{}
}

var DeployProxyCmd = &cobra.Command{
	Use:   "kube-proxy",
	Short: "kube-proxy",
	Long:  "kube-proxy",
	Run: func(cmd *cobra.Command, args []string) {

		conf := config.GetKubernetesConf()

		utils.DoOrDieWithMsg(utils.ExecShell([]string{common.GetCertShellPath(common.KubeProxy)},
			nil, nil, nil), "Failed to generate the kube-proxy cert")

		utils.DoOrDieWithMsg(newKubeProxyService().startService(conf),
			"Failed to start the kube-proxy service")
	},
}

func (kps *kubeProxyService) startService(conf *config.KubernetesConf) error {
	startArgs := []string{
		common.GetDeployShellPath(common.KubeProxy),
		conf.ClusterCIDR,
		fmt.Sprintf("https://%s:6443", conf.KubeApiServer),
	}
	return utils.ExecShell(startArgs, nil, nil, nil)
}
