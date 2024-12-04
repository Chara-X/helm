package action

import (
	"regexp"
	"strings"

	"github.com/Chara-X/util/slices"

	"github.com/Chara-X/helm"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

var sep = regexp.MustCompile("(?:^|\\s*\n)---\\s*")

type Uninstall struct {
	u   *action.Uninstall
	cfg *action.Configuration
}

func NewUninstall(cfg *action.Configuration) *Uninstall {
	if helm.Reference {
		return &Uninstall{u: action.NewUninstall(cfg)}
	}
	return &Uninstall{cfg: cfg}
}
func (u *Uninstall) Run(name string) (*release.UninstallReleaseResponse, error) {
	if helm.Reference {
		return u.u.Run(name)
	}
	var rls, _ = u.cfg.Releases.Driver.Get(name)
	var resources, _ = u.cfg.KubeClient.Build(strings.NewReader(strings.Join(slices.Reverse(sep.Split(rls.Manifest, -1)), "\n---\n")), false)
	u.cfg.KubeClient.Delete(resources)
	u.cfg.Releases.Driver.Delete(rls.Name)
	return &release.UninstallReleaseResponse{Release: rls}, nil
}
