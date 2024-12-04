package getter

import (
	"bytes"

	"github.com/Chara-X/helm"
	"helm.sh/helm/v3/pkg/getter"
)

type Getter interface {
	Get(url string) (*bytes.Buffer, error)
}

func NewHTTPGetter() (Getter, error) {
	if helm.Reference {
		var g, err = getter.NewHTTPGetter()
		return &httpGetter{g: g}, err
	}
	panic("unimplemented")
}

type httpGetter struct {
	g getter.Getter
}

func (g *httpGetter) Get(url string) (*bytes.Buffer, error) {
	if helm.Reference {
		return g.g.Get(url)
	}
	panic("unimplemented")
}
