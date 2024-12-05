package getter

import (
	"bytes"

	"helm.sh/helm/v3/pkg/getter"
)

type Getter interface {
	Get(url string) (*bytes.Buffer, error)
}

func NewHTTPGetter() (Getter, error) {
	var g, err = getter.NewHTTPGetter()
	return &httpGetter{g: g}, err
}

type httpGetter struct{ g getter.Getter }

func (g *httpGetter) Get(url string) (*bytes.Buffer, error) { return g.g.Get(url) }
