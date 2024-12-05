package chart

import (
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

type Chart struct {
	Metadata  *chart.Metadata        `json:"metadata"`
	Templates []*chart.File          `json:"templates"`
	Values    map[string]interface{} `json:"values"`
}

func Load(name string) (*Chart, error) {
	var c, err = loader.Load(name)
	return &Chart{Metadata: c.Metadata, Templates: c.Templates, Values: c.Values}, err
}
