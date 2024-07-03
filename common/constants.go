package common

type ProviderType string

const ProviderTypeKind ProviderType = "Kind"

type CNIType string

const (
	CNITypeCalico  CNIType = "Calico"
	CNITypeFlannel CNIType = "Flannel"
	CNITypeKindNet CNIType = "Kindnet"
	CNITypeCanal   CNIType = "Canal"
	CNITypeWave    CNIType = "Wave"
	CNITypeCilium  CNIType = "Cilium"
)
