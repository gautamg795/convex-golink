// Copyright 2022 Tailscale Inc & Contributors
// SPDX-License-Identifier: BSD-3-Clause

package golink

import (
	_ "embed"
	"net/url"
	"strings"
	"time"
)

// Link is the structure stored for each go short link.
type Link struct {
	Short    string // the "foo" part of http://go/foo
	Long     string // the target URL or text/template pattern to run
	Created  time.Time
	LastEdit time.Time // when the link was last edited
	Owner    string    // user@domain
}

// ClickStats is the number of clicks a set of links have received in a given
// time period. It is keyed by link short name, with values of total clicks.
type ClickStats map[string]int

// linkID returns the normalized ID for a link short name.
func linkID(short string) string {
	id := url.PathEscape(strings.ToLower(short))
	id = strings.ReplaceAll(id, "-", "")
	return id
}

type Database interface {
	LoadAll() ([]*Link, error)
	Load(short string) (*Link, error)
	Save(link *Link) error
	Delete(short string) error
	LoadStats() (ClickStats, error)
	SaveStats(stats ClickStats) error
	DeleteStats(short string) error
}
