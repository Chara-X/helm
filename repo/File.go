package repo

import (
	"helm.sh/helm/v3/pkg/repo"
)

type File struct {
	APIVersion   string        `json:"apiVersion"`
	Repositories []*repo.Entry `json:"repositories"`
}
