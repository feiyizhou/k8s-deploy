package etcd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s-deploy/common"
	"k8s-deploy/config"
	"k8s-deploy/utils"
	"strings"
)

type etcdService struct{}

func newEtcdService() *etcdService {
	return &etcdService{}
}

var DeployEtcdCmd = &cobra.Command{
	Use:   "etcd",
	Short: "etcd",
	Long:  "etcd",
	Run: func(cmd *cobra.Command, args []string) {

		etcdConf := config.GetEtcdConf()

		utils.DoOrDieWithMsg(newEtcdService().generateCert(etcdConf),
			"Generate etcd cert filed")

		utils.DoOrDieWithMsg(newEtcdService().startService(etcdConf),
			"Start etcd service failed")
	},
}

func (es *etcdService) generateCert(conf *config.EtcdConf) error {
	var hosts []string
	for _, node := range conf.Nodes {
		hosts = append(hosts, node.IP)
	}
	certArgs := []string{common.GetCertShellPath(common.ETCD)}
	certArgs = append(certArgs, hosts...)
	return utils.ExecShell(certArgs, nil, nil, nil)
}

func (es *etcdService) startService(conf *config.EtcdConf) error {
	var currentName, currentIP string
	currentName = conf.CurrentName
	var deployArgs []string
	if len(conf.Nodes) == 1 {
		deployArgs = []string{
			common.GetSingleDeployShellPath(common.ETCD),
			currentName,
		}
	} else {
		var clusterIPArr []string
		for _, node := range conf.Nodes {
			if strings.EqualFold(currentName, node.Name) {
				currentIP = node.IP
			}
			clusterIPArr = append(clusterIPArr,
				fmt.Sprintf("%s=https://%s:2380", node.Name, node.IP))
		}
		deployArgs = []string{
			common.GetDeployShellPath(common.ETCD),
			currentName,
			currentIP,
			strings.Join(clusterIPArr, ","),
		}
	}
	return utils.ExecShell(deployArgs, nil, nil, nil)
}
