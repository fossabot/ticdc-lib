package client

import (
	"context"
	"github.com/pingcap/ticdc-lib/config"
)

// AdminJobType represents for admin job type, both used in owner and processor
type AdminJobType int

// All AdminJob types
const (
	AdminNone AdminJobType = iota
	AdminStop
	AdminResume
	AdminRemove
)

// ChangeFeed used to display change feed info
type ChangeFeed struct {
	Name         string       `json:"name"`
	ResolvedTs   uint64       `json:"resolved-ts"`
	CheckpointTs uint64       `json:"checkpoint-ts"`
	AdminJobType AdminJobType `json:"admin-job-type"`
}

// CreateChangeFeed creates a change feed using the config
// returns the new change feed info and error
func (c *TiCDCClient) CreateChangeFeed(ctx context.Context, config *config.ChangeFeed) (*ChangeFeed, error) {
	panic("unimplemented")
}

// GetChangeFeedList returns a list of change feeds whose match the `name`,
// parameter `name` supports wildcard character
func (c *TiCDCClient) GetChangeFeedList(ctx context.Context, name string) ([]*ChangeFeed, error) {
	panic("unimplemented")
}

// PauseChangeFeed pauses the change feed named `name`
func (c *TiCDCClient) PauseChangeFeed(ctx context.Context, name string) error {
	panic("unimplemented")
}

// ResumeChangeFeed stops the change feed named `name`
func (c *TiCDCClient) ResumeChangeFeed(ctx context.Context, name string) error {
	panic("unimplemented")
}

// RemoveChangeFeed removes the change feed named `name`
func (c *TiCDCClient) RemoveChangeFeed(ctx context.Context, name string) (checkpointTS uint64, err error) {
	panic("unimplemented")
}
