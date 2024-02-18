package dracoon_go_sdk

import (
	"context"
	"fmt"
	"github.com/oliverziegert/dracoon-go-sdk/client/dracoon/core"
	cac "github.com/oliverziegert/dracoon-go-sdk/client/dracoon/core/openapi"

	"golang.org/x/oauth2"
)

type Client struct {
	cfg *core.Config
	ctx context.Context
	dc  *cac.APIClient
}

func NewCoreClient(cfg *core.Config, ts oauth2.TokenSource) *Client {

	if cfg == nil {
		cfg = core.NewDefaultConfig()
	}
	ctx := context.Background()
	dcCfg := cac.NewConfiguration()
	dcCfg.Servers = cac.ServerConfigurations{
		{
			URL:         fmt.Sprintf("https://%s/api", cfg.Domain),
			Description: "No description provided",
		},
	}
	dcCfg.OperationServers = map[string]cac.ServerConfigurations{}
	dcCfg.UserAgent = fmt.Sprintf("%s|%s", name, version)
	dcCfg.HTTPClient = cfg.Client

	client := &Client{
		ctx: context.WithValue(ctx, cac.ContextOAuth2, ts),
		cfg: cfg,
		dc:  cac.NewAPIClient(dcCfg),
	}
	return client
}
