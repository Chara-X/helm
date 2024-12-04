package search

import (
	"regexp"
	"strings"

	"github.com/Chara-X/helm"
	"github.com/Chara-X/helm/pkg/repo"
	"helm.sh/helm/v3/cmd/helm/search"
	repoRef "helm.sh/helm/v3/pkg/repo"
)

type Index struct {
	i      *search.Index
	charts []*repoRef.ChartVersion
}

func NewIndex() *Index {
	if helm.Reference {
		return &Index{i: search.NewIndex()}
	}
	return &Index{}
}
func (i *Index) AddRepo(ind *repo.IndexFile) {
	if helm.Reference {
		i.i.AddRepo("", &repoRef.IndexFile{APIVersion: ind.APIVersion, Entries: ind.Entries}, false)
		return
	}
	for _, charts := range ind.Entries {
		i.charts = append(i.charts, charts[0])
	}
}
func (i *Index) All() []*search.Result {
	if helm.Reference {
		return i.i.All()
	}
	var res = []*search.Result{}
	for _, chart := range i.charts {
		res = append(res, &search.Result{Chart: chart})
	}
	return res
}
func (i *Index) Search(term string, threshold int) ([]*search.Result, error) {
	if helm.Reference {
		return i.i.Search(term, threshold, true)
	}
	var res = []*search.Result{}
	for _, chart := range i.charts {
		if loc := regexp.MustCompile(strings.ToLower(term)).FindStringIndex(strings.ToLower(chart.Name + "\v" + chart.Description + "\v" + strings.Join(chart.Keywords, " "))); loc != nil {
			res = append(res, &search.Result{Score: loc[0], Chart: chart})
		}
	}
	return res, nil
}
