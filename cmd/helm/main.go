package main

import (
	"io"
	"os"
	"sync"

	"github.com/Chara-X/helm/getter"
	"github.com/Chara-X/helm/repo"
	"gopkg.in/yaml.v3"
	repoRef "helm.sh/helm/v3/pkg/repo"
)

var (
	repoConfig = os.Getenv("HELM_REPOSITORY_CONFIG")
	repoCache  = os.Getenv("HELM_REPOSITORY_CACHE")
)

func main() {
	switch os.Args[1] {
	case "repo":
		switch os.Args[2] {
		case "add":
			var f repo.File
			var data, _ = os.ReadFile(repoConfig)
			yaml.Unmarshal(data, &f)
			f.Repositories = append(f.Repositories, &repoRef.Entry{Name: os.Args[3], URL: os.Args[4]})
			data, _ = yaml.Marshal(f)
			os.WriteFile(repoConfig, data, 0600)
		case "update":
			var f repo.File
			var data, _ = os.ReadFile(repoConfig)
			yaml.Unmarshal(data, &f)
			var wg sync.WaitGroup
			for _, r := range f.Repositories {
				wg.Add(1)
				go func() {
					defer wg.Done()
					var cli, _ = getter.NewHTTPGetter()
					var resp, _ = cli.Get(r.URL + "/index.yaml")
					var cache, _ = os.Create(repoCache + "/" + r.Name + "-index.yaml")
					io.Copy(cache, resp)
				}()
			}
			wg.Wait()
		default:
			panic("unimplemented")
		}
	default:
		panic("unimplemented")
	}
}
