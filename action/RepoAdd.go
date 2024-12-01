package action

import (
	"os"

	"github.com/Chara-X/helm"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

type RepoAdd struct {
	Name string
	URL  string
}

func (r *RepoAdd) Run() error {
	var f repo.File
	var data, _ = os.ReadFile(helm.RepoConfig)
	yaml.Unmarshal(data, &f)
	var entry = repo.Entry{Name: r.Name, URL: r.URL}
	var repo, _ = repo.NewChartRepository(&entry, getter.All(cli.New()))
	repo.DownloadIndexFile()
	f.Update(&entry)
	return f.WriteFile(helm.RepoConfig, 0600)
}
