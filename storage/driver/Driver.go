package driver

import (
	"github.com/Chara-X/helm"
	rspb "helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Driver interface {
	Name() string
	Create(key string, rls *rspb.Release) error
	Update(key string, rls *rspb.Release) error
	Delete(key string) (*rspb.Release, error)
	Get(key string) (*rspb.Release, error)
	List(filter func(*rspb.Release) bool) ([]*rspb.Release, error)
	Query(labels map[string]string) ([]*rspb.Release, error)
}

func NewConfigMaps(impl core.ConfigMapInterface) Driver {
	if helm.Reference {
		return driver.NewConfigMaps(impl)
	}
	panic("unimplemented")
}
func NewSQL(connectionString string, namespace string) (Driver, error) {
	if helm.Reference {
		return driver.NewSQL(connectionString, func(s string, i ...interface{}) {}, namespace)
	}
	panic("unimplemented")
}
func NewMemory(ns string) Driver {
	if helm.Reference {
		var mem = driver.NewMemory()
		mem.SetNamespace(ns)
		return mem
	}
	panic("unimplemented")
}
