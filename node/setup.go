package node

import (
	"context"
	"dim-edge-core/protocol"

	"github.com/golang/protobuf/ptypes/empty"
)

// CheckSetup check if influx db has been setup
func (c *Client) CheckSetup() (setup *protocol.CheckSetupRes, err error) {
	setup, err = c.Store.CheckSetup(context.TODO(), &empty.Empty{})

	return
}

// Setup setup influxdb
func (c *Client) Setup(p *protocol.SetupParams) (err error) {
	_, err = c.Store.Setup(context.TODO(), p)

	return
}