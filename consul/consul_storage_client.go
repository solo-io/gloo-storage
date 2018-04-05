package consul

import (
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"github.com/solo-io/gloo-storage"
	"github.com/solo-io/gloo-storage/internal/base"
)

type Client struct {
	v1 *v1client
}

// TODO: support basic auth and tls
func NewStorage(cfg *api.Config, rootPath string, syncFrequency time.Duration) (storage.Interface, error) {
	cfg.WaitTime = syncFrequency

	// Get a new client
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "creating consul client")
	}

	return &Client{
		v1: &v1client{
			upstreams: &upstreamsClient{
				base: base.NewConsulStorageClient(rootPath+"/upstreams", client),
			},
			virtualHosts: &virtualHostsClient{
				base: base.NewConsulStorageClient(rootPath+"/virtualhosts", client),
			},
		},
	}, nil
}

func (c *Client) V1() storage.V1 {
	return c.v1
}

type v1client struct {
	upstreams    *upstreamsClient
	virtualHosts *virtualHostsClient
}

func (c *v1client) Register() error {
	return nil
}

func (c *v1client) Upstreams() storage.Upstreams {
	return c.upstreams
}

func (c *v1client) VirtualHosts() storage.VirtualHosts {
	return c.virtualHosts
}
