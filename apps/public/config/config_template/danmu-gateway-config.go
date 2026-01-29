package config_template

type DanmuGatewayConfig struct {
	Hertz Hertz
	Etcd  Etcd
}

type Hertz struct {
	ListenAddr     string
	ListenPort     string
	MonitoringPort string
}

type Etcd struct {
	ServiceName string
	Namespace   string
	Urls        []string
}
