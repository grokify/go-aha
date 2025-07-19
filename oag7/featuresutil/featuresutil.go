package featuresutil

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/grokify/mogo/encoding/jsonutil/jsonraw"
	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/net/urlutil"
)

func GetFeatureRaw(ctx context.Context, hclt *http.Client, baseURL, featureID string, prettyPrint bool) ([]byte, error) {
	featureID = strings.TrimSpace(featureID)
	sr := httpsimple.Request{
		Method: http.MethodGet,
		URL:    urlutil.JoinAbsolute(baseURL, "api/v1/features", featureID),
	}
	if hclt == nil {
		hclt = &http.Client{}
	}
	sc := httpsimple.NewClient(hclt, baseURL)
	resp, err := sc.Do(ctx, sr)
	if err != nil {
		return []byte{}, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	if prettyPrint {
		return jsonraw.IndentBytes(b, "", "  ")
	} else {
		return b, err
	}
}
