package client

import "context"

type Capture struct {
	ID      string `json:"id"`
	IsOwner bool   `json:"is-owner"`
}

type Processor struct {
	TableNames   []string     `json:"table-names"`
	AdminJobType AdminJobType `json:"admin-job-type"`
}

// GetCaptureList returns a list of captures whose match the `name`,
// parameter `captureID` supports wildcard character
func (c *TiCDCClient) GetCaptureList(ctx context.Context, captureID string) ([]*Capture, error) {
	panic("unimplemented")
}

// GetCaptureList returns the info of processor
func (c *TiCDCClient) GetProcessor(ctx context.Context, captureID string, changefeedName string) (*Processor, error) {
	panic("unimplemented")
}

// RetireOwner makes the owner retired
func (c *TiCDCClient) RetireOwner(ctx context.Context) error {
	panic("unimplemented")
}
