package action

import (
	"bytes"
	"text/template"

	"github.com/Chara-X/util/maps"

	"github.com/Chara-X/helm"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/cli-runtime/pkg/resource"
)

type Install struct {
	i           *action.Install
	cfg         *action.Configuration
	ReleaseName string
	Namespace   string
}

func NewInstall(cfg *action.Configuration) *Install {
	if helm.Reference {
		return &Install{i: action.NewInstall(cfg)}
	}
	return &Install{cfg: cfg}
}
func (i *Install) Run(chrt *chart.Chart, vals map[string]interface{}) (*release.Release, error) {
	if helm.Reference {
		i.i.ReleaseName, i.i.Namespace = i.ReleaseName, i.Namespace
		return i.i.Run(chrt, vals)
	}
	var buf = bytes.NewBuffer(nil)
	template.Must(template.New(chrt.Templates[0].Name).Parse(string(chrt.Templates[0].Data))).Execute(buf, map[string]interface{}{"Chart": chrt.Metadata, "Release": map[string]interface{}{"Name": i.ReleaseName, "Namespace": i.Namespace}, "Values": maps.Merge(chrt.Values, vals)})
	var resources, _ = resource.NewBuilder(i.cfg.RESTClientGetter).NamespaceParam(i.Namespace).Unstructured().Stream(buf, "").Do().Infos()
	i.cfg.KubeClient.Create(resources)
	var rls = &release.Release{Manifest: buf.String()}
	i.cfg.Releases.Driver.Create(i.ReleaseName, rls)
	return rls, nil
}
