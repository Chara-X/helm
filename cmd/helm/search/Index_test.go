package search

import (
	"fmt"

	"github.com/Chara-X/helm/pkg/repo"
	"helm.sh/helm/v3/pkg/chart"
	repoRef "helm.sh/helm/v3/pkg/repo"
)

func ExampleIndex() {
	var i = NewIndex()
	i.AddRepo(&repo.IndexFile{APIVersion: "v1", Entries: map[string]repoRef.ChartVersions{
		"nginx": {
			&repoRef.ChartVersion{Metadata: &chart.Metadata{
				Name:        "nginx",
				Description: "A high performance web server and a reverse proxy server",
				Keywords:    []string{"web", "server", "reverse", "proxy"},
			}},
		},
	}})
	var res, _ = i.Search("proxy", 0)
	for _, r := range res {
		fmt.Println(r.Score, r.Chart.Name)
	}
	// Output:
	// 50 nginx
}
