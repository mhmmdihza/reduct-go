package integration

import (
	"net/http"
	urlParser "net/url"
	"strconv"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mhmmdihza/reduct-go/reduct/integration/client"
)

type Integration struct {
	clientService *client.Reduct
}

type ClientOptions struct {
	ApiToken  string        // API token for authentication
	Timeout   time.Duration // communication timeout
	VerifySSL bool          // verify SSL certificate
}

func NewIntegration(url string, options *ClientOptions) (*Integration, error) {
	u, err := urlParser.Parse(url)
	if err != nil {
		return nil, err
	}
	cl := http.DefaultClient
	newTransport := func(httpClient *http.Client) *httptransport.Runtime {
		httpClient.Transport = &CustomTransportForceContentLengthHeader{
			Base: http.DefaultTransport,
		}
		tp := httptransport.NewWithClient(u.Host, "", []string{u.Scheme}, httpClient)
		return tp
	}

	if options == nil {
		return &Integration{client.New(newTransport(cl), strfmt.Default)}, nil
	}

	if options.VerifySSL {
		if cl, err = httptransport.TLSClient(httptransport.TLSClientOptions{
			InsecureSkipVerify: !options.VerifySSL,
		}); err != nil {
			return nil, err
		}
	}
	if options.Timeout > 0 {
		cl.Timeout = options.Timeout
	}

	transport := newTransport(cl)
	if options.ApiToken != "" {
		transport.DefaultAuthentication = httptransport.BearerToken(options.ApiToken)
	}

	return &Integration{client.New(transport, strfmt.Default)}, nil
}

type CustomTransportForceContentLengthHeader struct {
	Base http.RoundTripper
}

func (t *CustomTransportForceContentLengthHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	if v, ok := req.Header["Content-Length"]; ok && len(v) > 0 {
		req.ContentLength, _ = strconv.ParseInt(v[0], 10, 64)
	}
	transport := t.Base
	if transport == nil {
		transport = http.DefaultTransport
	}
	return transport.RoundTrip(req)
}
