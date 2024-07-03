package kind

import (
	appv1 "example.org/multi-clusters/api/app/v1"
)

type Kind struct {
	Host     string
	Port     int
	User     string
	Password string
}

func NewKindProvider() *Kind {
	return &Kind{}
}

func (k *Kind) ListClusters() {}

func (k *Kind) GetCluster(clusterName string) (interface{}, error) {
	return "", nil
}

func (k *Kind) CreateCluster(cluster *appv1.Cluster) error {
	return nil
}
func (k *Kind) UpdateCluster() {}
func (k *Kind) DeleteCluster() {}
func (k *Kind) StartCluster()  {}
func (k *Kind) StopCluster()   {}
