package helm

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
)

type Template struct {
	i           *action.Install
	ReleaseName string
	Namespace   string
}

func NewTemplate(cfg *action.Configuration) *Template {
	if Reference {
		return &Template{i: action.NewInstall(cfg)}
	}
	panic("not implemented")
}
func (t *Template) Run(chrt *chart.Chart, vals map[string]interface{}) (*release.Release, error) {
	if Reference {
		t.i.ReleaseName = t.ReleaseName
		t.i.Namespace = t.Namespace
		t.i.DryRun = true
		t.i.ClientOnly = true
		return t.i.Run(chrt, vals)
	}
	panic("not implemented")
}
