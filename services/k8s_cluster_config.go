package services

import (
	"log"
	"net/http"
	"time"

	v1 "github.com/SwaDeshiTech/arsenal/pkg/mongo-connector/v1"
	customerror "github.com/SwaDeshiTech/kubesync/customError"
	"github.com/SwaDeshiTech/kubesync/dto"
	"github.com/SwaDeshiTech/kubesync/entity"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddK8sClusterConfig(k8sClusterConfigRequest dto.K8sClusterConfigRequest) (dto.K8sClusterConfigResponse, customerror.Http) {

	var k8sClusterConfig entity.K8sClusterConfig
	err := mapstructure.Decode(k8sClusterConfigRequest, &k8sClusterConfig)
	if err != nil {
		log.Println("failed to map the dto & entity", err)
		return dto.K8sClusterConfigResponse{}, customerror.HttpError("Failed to map the DTO & Entity", http.StatusBadRequest)
	}

	k8sClusterConfig.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	k8sClusterConfig.IsActive = true

	err = k8sClusterConfig.Insert()
	if err != nil {
		log.Println("failed to insert k8sCluster config", err)
		return dto.K8sClusterConfigResponse{}, customerror.HttpError("Failed to onbaord K8s Cluster Config", http.StatusInternalServerError)
	}

	var k8sClusterConfigResponse dto.K8sClusterConfigResponse
	err = mapstructure.Decode(k8sClusterConfig, &k8sClusterConfigResponse)
	if err != nil {
		log.Println("Failed to map Entity & DTO", err)
		return dto.K8sClusterConfigResponse{}, customerror.HttpError("Failed to map Entity & DTO", http.StatusInternalServerError)
	}

	return k8sClusterConfigResponse, customerror.Http{}
}

func GetK8sClusterConfigDetail(k8sClusterName string) (dto.K8sClusterConfigResponse, customerror.Http) {

	filters := bson.M{"name": k8sClusterName}
	sort := bson.D{{Key: "name", Value: 1}}

	resultCriteria := v1.ResultCriteria{
		Filters: filters,
		Sort:    sort,
	}

	var k8sClusterConfigResponse dto.K8sClusterConfigResponse

	k8sClusterConfigs, err := entity.FetchK8sClusterConfigs(resultCriteria)
	if err != nil {
		log.Println("failed to fetch k8s cluster config", err)
		return k8sClusterConfigResponse, customerror.HttpError("k8s cluster config could not be fetched", http.StatusInternalServerError)
	}

	if len(k8sClusterConfigs) == 0 {
		log.Println("k8s cluster config could not be found")
		return k8sClusterConfigResponse, customerror.HttpError("k8s cluster config could not be found", http.StatusNotFound)
	}

	err = mapstructure.Decode(k8sClusterConfigs[0], &k8sClusterConfigResponse)
	if err != nil {
		log.Println("failed to map entity & dto", err)
		return k8sClusterConfigResponse, customerror.HttpError("k8s cluster config could not be fetched", http.StatusInternalServerError)
	}

	return k8sClusterConfigResponse, customerror.Http{}
}
