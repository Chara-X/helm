package driver

import (
	"github.com/Chara-X/helm"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Driver interface {
	Name() string
	Create(key string, rls *release.Release) error
	Update(key string, rls *release.Release) error
	Delete(key string) (*release.Release, error)
	Get(key string) (*release.Release, error)
	List(filter func(*release.Release) bool) ([]*release.Release, error)
	Query(labels map[string]string) ([]*release.Release, error)
}

func NewConfigMaps(impl core.ConfigMapInterface) Driver {
	if helm.Reference {
		return driver.NewConfigMaps(impl)
	}
	panic("unimplemented")
}
