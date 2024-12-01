package helm

import "os"

var Reference = false
var (
	RepoConfig = os.Getenv("HELM_REPOSITORY_CONFIG")
	RepoCache  = os.Getenv("HELM_REPOSITORY_CACHE")
)
