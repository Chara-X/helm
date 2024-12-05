package driver

import (
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func ExampleDriver() {
	var client, _ = kubernetes.NewForConfig(config.GetConfigOrDie())
	NewConfigMaps(client.CoreV1().ConfigMaps("default")).Create("example", &release.Release{Name: "example", Manifest: "apiVersion: v1", Info: &release.Info{}})
	// Output:
}
