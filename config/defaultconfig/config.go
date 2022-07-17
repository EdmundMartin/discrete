package defaultconfig

import (
	"github.com/EdmundMartin/discrete/config"
	"github.com/EdmundMartin/discrete/protocol"
)

type DefaultConfig struct {
	ScrapingSupport            bool
	ScrapingWithoutHashes      bool
	BlacklistFunction          func(client string) bool
	AllowAutomaticRegistration bool
	Interval                   protocol.TrackerInterval
	PrivateOnly                bool
}

func (d DefaultConfig) SupportScraping() bool {
	return d.ScrapingSupport
}

func (d DefaultConfig) AllowScrapingWithoutInfoHashes() bool {
	return d.ScrapingWithoutHashes
}

func (d DefaultConfig) BlacklistedClient(client string) bool {
	return d.BlacklistFunction(client)
}

func (d DefaultConfig) AutoRegister() bool {
	return d.AllowAutomaticRegistration
}

func (d DefaultConfig) TrackerInterval() protocol.TrackerInterval {
	return d.Interval
}

func (d DefaultConfig) IsPrivateOnly() bool {
	return d.PrivateOnly
}

func NewDefaultConfig(opts ...func(defaultConfig *DefaultConfig)) config.ConfigStore {
	def := &DefaultConfig{
		BlacklistFunction: func(client string) bool {
			return false
		},
		ScrapingSupport:            false,
		ScrapingWithoutHashes:      false,
		AllowAutomaticRegistration: true,
		Interval: protocol.TrackerInterval{
			MinIntervalSeconds:     180,
			DefaultIntervalSeconds: 180,
		},
		PrivateOnly: false,
	}
	for _, opt := range opts {
		opt(def)
	}
	return def
}
