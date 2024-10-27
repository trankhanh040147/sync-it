package common

import "time"

type Configuration struct {
	SourcePath      string        `env:"source_path" json:"source_path" envDefault:""`
	DestinationPath string        `env:"destination_path" json:"destination_path" envDefault:""`
	ExcludePatterns string        `env:"exclude_patterns" json:"exclude_patterns" envDefault:"" envSeparator:","`
	SyncInterval    time.Duration `env:"sync_interval" json:"sync_interval" envDefault:"24h"`
}
