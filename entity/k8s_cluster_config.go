package entity

import (
	"context"

	jsonV1 "github.com/SwaDeshiTech/arsenal/pkg/json/v1"
	v1 "github.com/SwaDeshiTech/arsenal/pkg/mongo-connector/v1"
	"github.com/SwaDeshiTech/kubesync/constants"
	"github.com/SwaDeshiTech/kubesync/databases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type K8sClusterConfig struct {
	ID                   string             `json:"id" bson:"id"`
	Name                 string             `json:"name" bson:"name"`
	DisplayName          string             `json:"displayName" bson:"displayName"`
	EndPoint             string             `json:"endpoint" bson:"endpoint"`
	ServerID             string             `json:"serverID" bson:"serverID"`
	ClientID             string             `json:"clientID" bson:"clientID"`
	TenantID             string             `json:"tenantID" bson:"tenantID"`
	KubeServerIP         string             `json:"kubeServerIP" bson:"kubeServerID"`
	CertificateAuthority string             `json:"certificateAuthority" bson:"certificateAuthority"`
	IsActive             bool               `json:"isActive" bson:"isActive"`
	CreatedAt            primitive.DateTime `json:"createdAt" bson:"createdAt"`
}

func (k8sClusterConfig *K8sClusterConfig) Insert() error {

	mongoClient := databases.MongoClient

	db := mongoClient.Database(constants.MONGO_KUBE_SYNC_DB_NAME)
	k8sClusterConfigCollection := db.Collection(constants.K8S_CLUSTER_CONFIGS_COLLECTION_NAME)

	_, err := k8sClusterConfigCollection.InsertOne(context.TODO(), k8sClusterConfig)
	if err != nil {
		return err
	}
	return nil
}

func FetchK8sClusterConfigs(resultCriteria v1.ResultCriteria) ([]K8sClusterConfig, error) {

	mongoClient := databases.MongoClient

	db := mongoClient.Database(constants.MONGO_KUBE_SYNC_DB_NAME)
	k8sClusterConfigCollection := db.Collection(constants.K8S_CLUSTER_CONFIGS_COLLECTION_NAME)

	if resultCriteria.Sort == nil {
		resultCriteria.Sort = primitive.D{}
	}

	findOptions := options.Find()
	findOptions.SetSort(resultCriteria.Sort)

	cursor, err := k8sClusterConfigCollection.Find(context.TODO(), resultCriteria.Filters, findOptions)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	jsonResp, err := jsonV1.ConvertIntoJSON(results)
	if err != nil {
		return nil, err
	}

	k8sClusterConfigList := []K8sClusterConfig{}

	err = jsonV1.ParseJSON(jsonResp, &k8sClusterConfigList)
	if err != nil {
		return nil, err
	}

	return k8sClusterConfigList, nil
}
