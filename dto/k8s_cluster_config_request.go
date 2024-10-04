package dto

type K8sClusterConfigRequest struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	DisplayName          string `json:"displayName"`
	EndPoint             string `json:"endpoint"`
	ServerID             string `json:"serverID"`
	ClientID             string `json:"clientID"`
	TenantID             string `json:"tenantID"`
	KubeServerIP         string `json:"kubeServerIP"`
	CertificateAuthority string `json:"certificateAuthority"`
}

type K8sClusterConfigParam struct {
	Name string `uri:"name,required"`
}
