package action

import (
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func ExampleUninstall() {
	var cfg = new(action.Configuration)
	cfg.Init(cli.New().RESTClientGetter(), "helm", "configmap", log.Printf)
	NewUninstall(cfg).Run("example")
	// Output:
}
