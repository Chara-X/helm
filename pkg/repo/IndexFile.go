package repo

import (
	"helm.sh/helm/v3/pkg/repo"
)

type IndexFile struct {
	APIVersion string                        `json:"apiVersion"`
	Entries    map[string]repo.ChartVersions `json:"entries"`
}

func IndexDirectory(dir, baseURL string) (*IndexFile, error) {
	var i, err = repo.IndexDirectory(dir, baseURL)
	return &IndexFile{APIVersion: i.APIVersion, Entries: i.Entries}, err
}
