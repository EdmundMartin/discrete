package config

import "github.com/EdmundMartin/discrete/protocol"

type ConfigStore interface {
	SupportScraping() bool
	AllowScrapingWithoutInfoHashes() bool
	BlacklistedClient(client string) bool
	AutoRegister() bool
	TrackerInterval() protocol.TrackerInterval
	IsPrivateOnly() bool
}
