package torrents

import "time"

type TorrentInfo struct {
	InfoHash  string
	Title     string
	CreatedOn time.Time
	UpdatedOn time.Time
}
