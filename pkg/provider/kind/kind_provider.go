package kind

import (
	"bytes"
	"errors"
	appv1 "example.org/multi-clusters/api/app/v1"
	"example.org/multi-clusters/common"
	"github.com/go-logr/logr"
	"os/exec"
	"strings"
)

type Kind struct {
	Log      logr.Logger
	Host     string
	Port     int
	User     string
	Password string
}

func NewKindProvider(logger logr.Logger) *Kind {
	return &Kind{
		Log: logger,
	}
}

func (k *Kind) ListClusters() ([]string, error) {
	output, err := RunKindCommand("get", "clusters")
	if err != nil {
		k.Log.Error(err, output)
		if errors.Is(err, exec.ErrNotFound) {
			// command not found
			return nil, err
		}
		// other error
		return nil, err
	}
	clusterList := strings.Split(strings.TrimSpace(output), "\n")
	return clusterList, nil
}

func (k *Kind) GetCluster(clusterName string) (interface{}, error) {
	output, err := RunKindCommand("get", "clusters")
	if err != nil {
		k.Log.Error(err, output)
		if errors.Is(err, exec.ErrNotFound) {
			// command not found
			return nil, err
		}
		// other error
		return nil, err
	}
	clusterList := strings.Split(strings.TrimSpace(output), "\n")
	for _, c := range clusterList {
		if c == clusterName {
			return "", nil
		}
	}
	return nil, common.ClusterNotFoundError
}

func (k *Kind) CreateCluster(cluster *appv1.Cluster) error {
	output, err := RunKindCommand("create", "cluster", "--name ", cluster.Name)
	if err != nil {
		k.Log.Error(err, output)
		if errors.Is(err, exec.ErrNotFound) {
			// command not found
			return err
		}
		// other error
		return err
	}
	return nil
}
func (k *Kind) UpdateCluster() {}
func (k *Kind) DeleteCluster() {}
func (k *Kind) StartCluster()  {}
func (k *Kind) StopCluster()   {}

func RunKindCommand(commands ...string) (string, error) {
	cmd := exec.Command("kind", commands...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}
