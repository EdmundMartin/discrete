package config

type ConfigStore interface {
	SupportScraping() bool
	AllowScrapingWithoutInfoHashes() bool
	BlacklistedClient(client string) bool
	AutoRegister() bool
}
