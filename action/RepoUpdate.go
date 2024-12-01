package action

import (
	"sync"

	"github.com/Chara-X/helm"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

type RepoUpdate struct{}

func (r *RepoUpdate) Run() error {
	var f, _ = repo.LoadFile(helm.RepoConfig)
	var repos []*repo.ChartRepository
	for _, cfg := range f.Repositories {
		var repo, _ = repo.NewChartRepository(cfg, getter.All(cli.New()))
		repos = append(repos, repo)
	}
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func() {
			defer wg.Done()
			repo.DownloadIndexFile()
		}()
	}
	wg.Wait()
	return nil
}
