package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type K8sClusterConfigResponse struct {
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	DisplayName          string             `json:"displayName"`
	EndPoint             string             `json:"endpoint"`
	ServerID             string             `json:"serverID"`
	ClientID             string             `json:"clientID"`
	TenantID             string             `json:"tenantID"`
	KubeServerIP         string             `json:"kubeServerIP"`
	CertificateAuthority string             `json:"certificateAuthority"`
	CreatedAt            primitive.DateTime `json:"createdAt"`
}
