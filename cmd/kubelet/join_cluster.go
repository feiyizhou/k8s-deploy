package kubelet

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"k8s-deploy/utils"
	"os"
	"os/exec"
	"strings"
)

type kubeletJoinClusterService struct{}

func newKubeletJoinClusterService() *kubeletJoinClusterService {
	return &kubeletJoinClusterService{}
}

var JoinClusterCmd = &cobra.Command{
	Use:   "join-cluster",
	Short: "join-cluster",
	Long:  "join-cluster",
	Run: func(cmd *cobra.Command, args []string) {
		pendingNodeName, err := newKubeletJoinClusterService().getUnJoinedNodeName()
		utils.CheckErr(err)

		if len(pendingNodeName) != 0 {
			err = newKubeletJoinClusterService().joinCluster(pendingNodeName)
			utils.CheckErr(err)
		}
	},
}

func (kjs *kubeletJoinClusterService) joinCluster(nodeNameArr []string) error {
	for _, nodeName := range nodeNameArr {
		args := []string{
			"certificate",
			"approve",
			nodeName,
		}
		cmd := exec.Command("kubectl", args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (kjs *kubeletJoinClusterService) getUnJoinedNodeName() ([]string, error) {
	var pendingNodeName []string
	file, err := os.OpenFile("csr.tmp", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cmd := exec.Command("kubectl", []string{"get", "csr"}...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = file
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	fileData, err := os.Open("csr.tmp")
	if err != nil {
		return nil, err
	}
	defer fileData.Close()
	br := bufio.NewReader(fileData)
	for {
		bytes, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tmpArr := strings.Split(string(bytes), " ")
		if strings.HasPrefix(tmpArr[0], "node-csr-") &&
			strings.EqualFold(tmpArr[len(tmpArr)-1], "Pending") {
			pendingNodeName = append(pendingNodeName, tmpArr[0])
		}
	}
	defer func() {
		err := os.Remove("csr.tmp")
		if err != nil {
			fmt.Println(err)
		}
	}()
	return pendingNodeName, err
}
