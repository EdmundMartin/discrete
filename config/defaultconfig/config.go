package defaultconfig

import "github.com/EdmundMartin/discrete/config"

type DefaultConfig struct {
	ScrapingSupport            bool
	ScrapingWithoutHashes      bool
	BlacklistFunction          func(client string) bool
	AllowAutomaticRegistration bool
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

func NewDefaultConfig(opts ...func(defaultConfig *DefaultConfig)) config.ConfigStore {
	def := &DefaultConfig{
		BlacklistFunction: func(client string) bool {
			return false
		},
		ScrapingSupport:            false,
		ScrapingWithoutHashes:      false,
		AllowAutomaticRegistration: true,
	}
	for _, opt := range opts {
		opt(def)
	}
	return def
}
