package config

import (
	"LiveDanmu/apps/public/config/config_reader"
	"LiveDanmu/apps/public/config/config_template"
	"LiveDanmu/apps/public/config/dns_lookup"
)

func LoadDanmuGatewayConfig() (*config_template.DanmuGatewayConfig, error) {
	// 初始化配置文件主体
	conf, err := config_reader.DanmuGatewayConfigLoader()
	if err != nil {
		return nil, err
	}
	// 服务发现
	addrList, err := dns_lookup.ServiceDiscovery(conf.Etcd.ServiceName, conf.Etcd.Namespace)
	if err != nil {
		return nil, err
	}
	// 组装配置
	conf.Etcd.Urls = addrList

	return conf, nil
}
