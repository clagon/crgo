package model

import (
	"encoding/json"
	"testing"
)

func TestJsonLocalizedNameDecodesString(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantName JsonLocalizedName
	}{
		{name: "string_name", body: `{"name":"Arena 17"}`, wantName: "Arena 17"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var arena Arena
			if err := json.Unmarshal([]byte(tt.body), &arena); err != nil {
				t.Fatal(err)
			}
			if arena.Name != tt.wantName {
				t.Fatalf("arena name: got %q, want %q", arena.Name, tt.wantName)
			}
		})
	}
}

func TestFloatDecodesNumber(t *testing.T) {
	tests := []struct {
		name string
		body string
		want Float
	}{
		{name: "decimal", body: `1.5`, want: 1.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var value Float
			if err := json.Unmarshal([]byte(tt.body), &value); err != nil {
				t.Fatal(err)
			}
			if value != tt.want {
				t.Fatalf("value: got %v, want %v", value, tt.want)
			}
		})
	}
}
