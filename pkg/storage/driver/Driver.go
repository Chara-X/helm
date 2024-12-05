package driver

import (
	// "context"

	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/Chara-X/helm"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedCore "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Driver interface {
	Name() string
	driver.Creator
	driver.Updator
	driver.Deletor
	driver.Queryor
}
type configMaps struct {
	cfgmaps driver.Driver
	impl    typedCore.ConfigMapInterface
}

func NewConfigMaps(impl typedCore.ConfigMapInterface) Driver {
	if helm.Reference {
		return &configMaps{cfgmaps: driver.NewConfigMaps(impl)}
	}
	return &configMaps{impl: impl}
}
func (c *configMaps) Name() string {
	if helm.Reference {
		return c.cfgmaps.Name()
	}
	return "ConfigMap"
}
func (c *configMaps) Create(key string, rls *release.Release) error {
	if helm.Reference {
		return c.cfgmaps.Create(key, rls)
	}
	var buf bytes.Buffer
	var w = gzip.NewWriter(&buf)
	json.NewEncoder(w).Encode(rls)
	w.Close()
	c.impl.Create(context.Background(), &core.ConfigMap{ObjectMeta: meta.ObjectMeta{Name: key, Labels: map[string]string{"owner": "helm"}}, Data: map[string]string{"release": base64.StdEncoding.EncodeToString(buf.Bytes())}}, meta.CreateOptions{})
	return nil
}
func (c *configMaps) Update(key string, rls *release.Release) error {
	return c.cfgmaps.Update(key, rls)
}
func (c *configMaps) Delete(key string) (*release.Release, error) {
	return c.cfgmaps.Delete(key)
}
func (c *configMaps) Get(key string) (*release.Release, error) {
	return c.cfgmaps.Get(key)
}
func (c *configMaps) List(filter func(*release.Release) bool) ([]*release.Release, error) {
	return c.cfgmaps.List(filter)
}
func (c *configMaps) Query(labels map[string]string) ([]*release.Release, error) {
	return c.cfgmaps.Query(labels)
}
