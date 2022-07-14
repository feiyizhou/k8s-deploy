package kubelet

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
)

type kubeletStartService struct{}

func newKubeletStartService() *kubeletStartService {
	return &kubeletStartService{}
}

var DeployKubeletCmd = &cobra.Command{
	Use:   "kubelet",
	Short: "kubelet",
	Long:  "kubelet",
	Run: func(cmd *cobra.Command, args []string) {

		conf := config.GetKubernetesConf()

		utils.DoOrDieWithMsg(newKubeletStartService().startService(conf),
			"Failed to start kubelet service")

	},
}

func (kls *kubeletStartService) startService(conf *config.KubernetesConf) error {
	startArgs := []string{
		common.GetDeployShellPath(common.Kubelet),
		conf.HostnameOverride,
		conf.ClusterDNS,
		fmt.Sprintf("https://%s:6443", conf.KubeApiServer),
	}
	return utils.ExecShell(startArgs, nil, nil, nil)
}
