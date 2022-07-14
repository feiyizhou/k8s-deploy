package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"k8s-deploy/common"
	"k8s-deploy/utils"
	"reflect"
)

type EtcdConf struct {
	CurrentName string `json:"currentName" mapstructure:"currentName"`
	Nodes       []Node `json:"nodes"  mapstructure:"nodes"`
}

type Node struct {
	Name string `json:"name"  mapstructure:"name"`
	IP   string `json:"ip"  mapstructure:"ip"`
}

type KubernetesConf struct {
	HostnameOverride      string   `json:"hostnameOverride" mapstructure:"hostnameOverride"`
	AdvertiseAddress      string   `json:"advertiseAddress"  mapstructure:"advertiseAddress"`
	ServerHosts           []string `json:"serverHosts"  mapstructure:"serverHosts"`
	EtcdHosts             []string `json:"etcdHosts"  mapstructure:"etcdHosts"`
	ServiceClusterIPRange string   `json:"serviceClusterIPRange"  mapstructure:"serviceClusterIPRange"`
	ClusterCIDR           string   `json:"clusterCIDR" mapstructure:"clusterCIDR"`
	KubeApiServer         string   `json:"kubeApiServer"  mapstructure:"kubeApiServer"`
	ClusterDNS            string   `json:"clusterDNS" mapstructure:"clusterDNS"`
}

type FlannelConf struct {
	Network          string `json:"network" mapstructure:"network"`
	Type             string `json:"type" mapstruture:"type"`
	KubeApiServer    string `json:"kubeApiServer"  mapstructure:"kubeApiServer"`
	HostnameOverride string `json:"hostnameOverride" mapstructure:"hostnameOverride"`
	IfaceName        string `json:"ifaceName" mapsructure:"ifaceName"`
}

func GetEtcdConf() *EtcdConf {
	conf, err := initEtcdConf()
	utils.CheckErr(err)
	utils.DieWithMsg(conf == nil, "Etcd configuration is not exist")
	return conf
}

func GetKubernetesConf() *KubernetesConf {
	conf, err := initKubernetesConf()
	utils.CheckErr(err)
	utils.DieWithMsg(conf == nil, "Kubernetes configuration is not exist")
	return conf
}

func GetFlannelConf() *FlannelConf {
	conf, err := initFlannelConf()
	utils.CheckErr(err)
	utils.DieWithMsg(conf == nil, "Flannel configuration is not exist")
	return conf
}

func initEtcdConf() (*EtcdConf, error) {
	viper.AddConfigPath(common.YamlConfigHomePath)
	viper.SetConfigName(common.ETCD)
	viper.SetConfigType(common.YamlConfigType)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	etcdConfMap := viper.GetStringMap(common.ETCD)
	if etcdConfMap == nil {
		return nil, fmt.Errorf("Etcd config is not exist ")
	}
	etcd := new(EtcdConf)
	err = ParseInterface2Struct(etcdConfMap, &etcd)
	if err != nil {
		return nil, fmt.Errorf("Parse confStr : %+v to struct err , err : %+v ", etcdConfMap, err)
	}
	return etcd, err
}

func initKubernetesConf() (*KubernetesConf, error) {
	viper.AddConfigPath(common.YamlConfigHomePath)
	viper.SetConfigName(common.Kubernetes)
	viper.SetConfigType(common.YamlConfigType)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	kubernetesConfMap := viper.GetStringMap(common.Kubernetes)
	if kubernetesConfMap == nil {
		return nil, fmt.Errorf("Kubernetes config is not exist ")
	}
	kubernetes := new(KubernetesConf)
	err = ParseInterface2Struct(kubernetesConfMap, kubernetes)
	if err != nil {
		return nil, fmt.Errorf("Parse confStr : %+v to struct err , err : %+v ", kubernetesConfMap, err)
	}
	return kubernetes, err
}

func initFlannelConf() (*FlannelConf, error) {
	viper.AddConfigPath(common.YamlConfigHomePath)
	viper.SetConfigName(common.Flannel)
	viper.SetConfigType(common.YamlConfigType)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	flannelConfMap := viper.GetStringMap(common.Flannel)
	if flannelConfMap == nil {
		return nil, fmt.Errorf("Kubernetes config is not exist ")
	}
	flannel := new(FlannelConf)
	err = ParseInterface2Struct(flannelConfMap, flannel)
	if err != nil {
		return nil, fmt.Errorf("Parse confStr : %+v to struct err , err : %+v ", flannelConfMap, err)
	}
	return flannel, err
}

// ParseInterface2Struct ...
func ParseInterface2Struct(in interface{}, out interface{}) error {
	kind := reflect.TypeOf(in).Kind()
	if reflect.Map == kind {
		err := mapstructure.Decode(in, out)
		if err != nil {
			return err
		}
	} else if reflect.String == kind {
		err := json.Unmarshal([]byte(in.(string)), out)
		if err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("Can not parse this type : %s ! ", kind.String()))
	}
	return nil
}
