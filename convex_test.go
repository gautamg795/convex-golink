// Copyright 2022 Tailscale Inc & Contributors
// SPDX-License-Identifier: BSD-3-Clause

package golink

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/joho/godotenv"
)

func clear(c *ConvexDB) {
	c.mutation(&UdfExecution{Path: "clear", Args: map[string]interface{}{}, Format: "json"})
}

func getDbUrl() string {
	envLocal, err := godotenv.Read(".env.local")
	if err != nil {
		return "https://feeble-gull-946.convex.cloud"
	}
	url := envLocal["VITE_CONVEX_URL"]
	if len(url) == 0 {
		url = "https://feeble-gull-946.convex.cloud"
	}
	return url
}

// Test saving and loading links for SQLiteDB
func Test_Convex_SaveLoadDeleteLinks(t *testing.T) {
	url := getDbUrl()
	db := NewConvexDB(url, "test")
	clear(db)
	defer clear(db)

	links := []*Link{
		{Short: "short", Long: "long", Created: time.Now(), LastEdit: time.Now()},
		{Short: "Foo.Bar", Long: "long", Created: time.Now(), LastEdit: time.Now()},
	}

	approximateTimeEq := cmp.Options{cmp.FilterPath(func(path cmp.Path) bool { return path.Last().Type() == reflect.TypeOf(time.Time{}) }, cmpopts.EquateApproxTime(time.Second))}
	for _, link := range links {
		if err := db.Save(link); err != nil {
			t.Error(err)
		}
		got, err := db.Load(link.Short)
		if err != nil {
			t.Error(err)
		}
		if !cmp.Equal(got, link, approximateTimeEq) {
			t.Errorf("db save and load got %v, want %v", *got, *link)
		}
	}

	got, err := db.LoadAll()
	if err != nil {
		t.Error(err)
	}

	sortLinks := cmpopts.SortSlices(func(a, b *Link) bool {
		return a.Short < b.Short
	})
	if !cmp.Equal(got, links, sortLinks, approximateTimeEq) {
		t.Errorf("db.LoadAll got %v, want %v", got, links)
	}

	for _, link := range links {
		if err := db.Delete(link.Short); err != nil {
			t.Error(err)
		}
	}

	got, err = db.LoadAll()
	if err != nil {
		t.Error(err)
	}
	want := []*Link(nil)
	if !cmp.Equal(got, want) {
		t.Errorf("db.LoadAll got %v, want %v", got, want)
	}
}

// Test saving and loading stats for SQLiteDB
func Test_Convex_SaveLoadDeleteStats(t *testing.T) {
	url := getDbUrl()
	db := NewConvexDB(url, "test")
	clear(db)
	// defer clear(db)

	// preload some links
	links := []*Link{
		{Short: "a"},
		{Short: "B-c"},
	}
	for _, link := range links {
		if err := db.Save(link); err != nil {
			t.Error(err)
		}
	}

	// Stats to store do not need to be their canonical short name,
	// but returned stats always should be.
	stats := []ClickStats{
		{"a": 1},
		{"b-c": 1},
		{"a": 1, "bc": 2},
	}
	want := ClickStats{
		"a":   2,
		"B-c": 3,
	}

	for _, s := range stats {
		if err := db.SaveStats(s); err != nil {
			t.Error(err)
		}
	}

	got, err := db.LoadStats()
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("db.LoadStats got %v, want %v", got, want)
	}

	for k := range want {
		if err := db.DeleteStats(k); err != nil {
			t.Error(err)
		}
	}

	got, err = db.LoadStats()
	if err != nil {
		t.Error(err)
	}
	want = ClickStats{}
	if !cmp.Equal(got, want) {
		t.Errorf("db.LoadStats got %v, want %v", got, want)
	}
}
