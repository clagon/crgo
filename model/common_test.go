package model

import (
	"encoding/json"
	"testing"
)

func TestJsonLocalizedNameDecodesString(t *testing.T) {
	var arena Arena
	if err := json.Unmarshal([]byte(`{"name":"Arena 17"}`), &arena); err != nil {
		t.Fatal(err)
	}

	if arena.Name != "Arena 17" {
		t.Fatalf("unexpected arena name: %q", arena.Name)
	}
}
