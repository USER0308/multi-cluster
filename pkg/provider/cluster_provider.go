package provider

import (
	appv1 "example.org/multi-clusters/api/app/v1"
	"example.org/multi-clusters/common"
	"example.org/multi-clusters/pkg/provider/kind"
	"github.com/go-logr/logr"
)

type ClusterProvider interface {
	ListClusters()
	GetCluster(clusterName string) (cluster interface{}, err error)
	CreateCluster(cluster *appv1.Cluster) error
	UpdateCluster()
	DeleteCluster()
	StartCluster()
	StopCluster()
}

func GetClusterProvider(logger logr.Logger, providerType common.ProviderType) ClusterProvider {
	switch providerType {
	case common.ProviderTypeKind:
		return kind.NewKindProvider(logger)
	}
	// not supported yet
	return nil
}
