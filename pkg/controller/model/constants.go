package model

const (
	//Grafana
	GrafanaImage                        = "grafana/grafana"
	GrafanaVersion                      = "7.1.1"
	GrafanaServiceAccountName           = "grafana-serviceaccount"
	GrafanaServiceName                  = "grafana-service"
	GrafanaDataStorageName              = "grafana-pvc"
	GrafanaConfigName                   = "grafana-config"
	GrafanaConfigFileName               = "grafana.ini"
	GrafanaIngressName                  = "grafana-ingress"
	GrafanaRouteName                    = "grafana-route"
	GrafanaDeploymentName               = "grafana-deployment"
	GrafanaPluginsVolumeName            = "grafana-plugins"
	GrafanaInitContainerName            = "grafana-plugins-init"
	GrafanaLogsVolumeName               = "grafana-logs"
	GrafanaDataVolumeName               = "grafana-data"
	GrafanaDatasourcesConfigMapName     = "grafana-datasources"
	GrafanaHealthEndpoint               = "/api/health"
	GrafanaPodLabel                     = "grafana"
	LastConfigAnnotation                = "last-config"
	LastConfigEnvVar                    = "LAST_CONFIG"
	LastDatasourcesConfigEnvVar         = "LAST_DATASOURCES"
	GrafanaAdminSecretName              = "grafana-admin-credentials"
	DefaultAdminUser                    = "admin"
	GrafanaAdminUserEnvVar              = "GF_SECURITY_ADMIN_USER"
	GrafanaAdminPasswordEnvVar          = "GF_SECURITY_ADMIN_PASSWORD"
	GrafanaHttpPort                 int = 3000
	GrafanaHttpPortName                 = "grafana"
	//Loki
	LokiImage              = "grafana/loki"
	LokiVersion            = "2.0.0"
	LokiHttpPort       int = 3100
	LokiHttpPortName       = "loki"
	LokiPodLabel           = "loki"
	LokiDeploymentName     = "loki-deployment"
)
