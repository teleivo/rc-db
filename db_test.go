package db_test

import (
	"testing"

	db "github.com/teleivo/rc-pairing"
)

func TestDB(t *testing.T) {
	database := db.New()

	v, ok := database.Get("a")
	if ok || v != "" {
		t.Fatalf("Get(%q) = %q, %t; want = %q, %t", "a", v, ok, "", false)
	}

	database.Set("a", "1")

	v, ok = database.Get("a")
	if !ok || v != "1" {
		t.Fatalf("Get(%q) = %q, %t; want = %q, %t", "a", v, ok, "1", true)
	}

	database.Set("a", "2")

	v, ok = database.Get("a")
	if !ok || v != "2" {
		t.Fatalf("Get(%q) = %q, %t; want = %q, %t", "a", v, ok, "2", true)
	}
}
