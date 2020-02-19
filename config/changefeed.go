package config

import "github.com/pingcap/tidb-tools/pkg/filter"

// ChangeFeed includes the items of ChangeFeed config
type ChangeFeed struct {
	Name    string  `toml:"name",json:"name"`
	SinkUri string  `toml:"sink_uri",json:"sink_uri"`
	Filter  *Filter `toml:"filter",json:"filter"`
}

// Filter represents some addition replication config for a changefeed
type Filter struct {
	FilterCaseSensitive bool          `toml:"filter-case-sensitive" json:"filter-case-sensitive"`
	FilterRules         *filter.Rules `toml:"filter-rules" json:"filter-rules"`
	IgnoreTxnCommitTs   []uint64      `toml:"ignore-txn-commit-ts" json:"ignore-txn-commit-ts"`
}
