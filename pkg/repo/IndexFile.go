package repo

import (
	"github.com/Chara-X/helm"
	"helm.sh/helm/v3/pkg/repo"
)

type IndexFile struct {
	APIVersion string                        `json:"apiVersion"`
	Entries    map[string]repo.ChartVersions `json:"entries"`
}

func IndexDirectory(dir, baseURL string) (*IndexFile, error) {
	if helm.Reference {
		var i, err = repo.IndexDirectory(dir, baseURL)
		return &IndexFile{APIVersion: i.APIVersion, Entries: i.Entries}, err
	}
	panic("unimplemented")
}
