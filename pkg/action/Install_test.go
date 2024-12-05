package action

import (
	"fmt"
	"log"

	"github.com/Chara-X/helm/pkg/chart"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func ExampleInstall() {
	var cfg = new(action.Configuration)
	cfg.Init(cli.New().RESTClientGetter(), "helm", "configmap", log.Printf)
	var i = NewInstall(cfg)
	i.ReleaseName, i.Namespace = "example", "helm"
	var chrt, _ = chart.Load("example")
	var vals = map[string]interface{}{"ownerBy": "chara"}
	var rls, _ = i.Run(chrt, vals)
	fmt.Println(rls.Manifest)
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
