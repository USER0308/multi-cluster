package provider

import (
	appv1 "example.org/multi-clusters/api/app/v1"
	"example.org/multi-clusters/common"
	"example.org/multi-clusters/pkg/provider/kind"
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

func GetClusterProvider(providerType common.ProviderType) ClusterProvider {
	switch providerType {
	case common.ProviderTypeKind:
		return kind.NewKindProvider()
	}
	// not supported yet
	return nil
}
