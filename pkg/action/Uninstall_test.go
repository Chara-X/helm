package action

import (
	"fmt"
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func ExampleUninstall() {
	var cfg = new(action.Configuration)
	cfg.Init(cli.New().RESTClientGetter(), "helm", "configmap", log.Printf)
	var u = NewUninstall(cfg)
	var rls, _ = u.Run("example")
	fmt.Println(rls.Release.Manifest)
	// Output:
	// apiVersion: v1
	// kind: Namespace
	// metadata:
	//   name: helm
	// ---
	// apiVersion: v1
	// kind: Pod
	// metadata:
	//   name: example
	//   namespace: helm
	//   labels:
	//     ownerBy: chara
	// spec:
	//   containers:
	//   - name: container
	//     image: docker.io/bitnami/nginx:1.27.0-debian-12-r3
	//     ports:
	//     - containerPort: 80
}
