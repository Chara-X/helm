package main

import (
	"io"
	"os"

	"github.com/Chara-X/helm/pkg/getter"
	"github.com/Chara-X/helm/pkg/repo"
	"gopkg.in/yaml.v3"
	repoRef "helm.sh/helm/v3/pkg/repo"
)

var (
	repoConfig = os.Getenv("HELM_REPOSITORY_CONFIG")
	repoCache  = os.Getenv("HELM_REPOSITORY_CACHE")
)

func main() {
	switch os.Args[0] {
	case "repo":
		switch os.Args[1] {
		case "update":
			var repos repo.File
			var data, _ = os.ReadFile(repoConfig)
			yaml.Unmarshal(data, &repos)
			for _, r := range repos.Repositories {
				var get, _ = getter.NewHTTPGetter()
				var res, _ = get.Get(r.URL + "/index.yaml")
				var cache, _ = os.Create(repoCache + "/" + r.Name + "-index.yaml")
				io.Copy(cache, res)
			}
		case "add":
			var repos repo.File
			var data, _ = os.ReadFile(repoConfig)
			yaml.Unmarshal(data, &repos)
			repos.Repositories = append(repos.Repositories, &repoRef.Entry{Name: os.Args[3], URL: os.Args[4]})
			data, _ = yaml.Marshal(repos)
			os.WriteFile(repoConfig, data, 0600)
		}
	default:
		panic("unimplemented")
	}
}
