package integration

import (
	"net/http"
	urlParser "net/url"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mhmmdihza/reduct-go/reduct/integration/client/operations"
)

type Integration struct {
	clientService operations.ClientService
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
		return httptransport.NewWithClient(u.Host, "", []string{u.Scheme}, httpClient)
	}

	if options == nil {
		return &Integration{operations.New(newTransport(cl), strfmt.Default)}, nil
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

	return &Integration{operations.New(transport, strfmt.Default)}, nil
}
