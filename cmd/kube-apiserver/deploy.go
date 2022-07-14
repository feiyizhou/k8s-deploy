package kube_apiserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
	"strings"
)

type kubeApiServerService struct{}

func newKubeApiServerService() *kubeApiServerService {
	return &kubeApiServerService{}
}

var DeployApiServerCmd = &cobra.Command{
	Use:   "kube-apiserver",
	Short: "kube-apiserver",
	Long:  "kube-apiserver",
	Run: func(cmd *cobra.Command, args []string) {

		conf := config.GetKubernetesConf()

		utils.DoOrDieWithMsg(newKubeApiServerService().generate(conf),
			"Generate the cert of kube-apiserver failed")

		utils.DoOrDieWithMsg(newKubeApiServerService().startService(conf),
			"Start the kube-apiserver service failed")
	},
}

func (kas *kubeApiServerService) generate(conf *config.KubernetesConf) error {
	certArgs := []string{common.GetCertShellPath(common.KubeApiServer)}
	certArgs = append(certArgs, conf.ServerHosts...)
	return utils.ExecShell(certArgs, nil, nil, nil)
}

func (kas *kubeApiServerService) startService(conf *config.KubernetesConf) error {
	var etcdServersArr []string
	for _, etcdHost := range conf.EtcdHosts {
		etcdServersArr = append(etcdServersArr,
			fmt.Sprintf("https://%s:2379", etcdHost))
	}
	startArgs := []string{
		common.GetDeployShellPath(common.KubeApiServer),
		strings.Join(etcdServersArr, ","),
		conf.AdvertiseAddress,
		conf.ServiceClusterIPRange,
	}
	return utils.ExecShell(startArgs, nil, nil, nil)
}
