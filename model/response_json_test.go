package model

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestResponseJSONPreservesZeroValues(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{
			name:  "zero crowns",
			value: PlayerBattleData{Crowns: 0},
			want:  `"crowns":0`,
		},
		{
			name:  "false boolean",
			value: CancelMatchResponse{Success: false},
			want:  `"success":false`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.value)
			if err != nil {
				t.Fatalf("json.Marshal() error = %v", err)
			}

			if !bytes.Contains(got, []byte(tt.want)) {
				t.Errorf("json.Marshal() = %s, want field %s", got, tt.want)
			}
		})
	}
}

func TestItemElixirCostDistinguishesMissingAndZero(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantNil   bool
		wantValue int
	}{
		{
			name:    "missing elixir cost",
			input:   `{"name":"Mirror"}`,
			wantNil: true,
		},
		{
			name:      "zero elixir cost",
			input:     `{"name":"Example","elixirCost":0}`,
			wantValue: 0,
		},
		{
			name:      "positive elixir cost",
			input:     `{"name":"Knight","elixirCost":3}`,
			wantValue: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Item
			if err := json.Unmarshal([]byte(tt.input), &got); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}

			if tt.wantNil {
				if got.ElixirCost != nil {
					t.Errorf("ElixirCost = %v, want nil", *got.ElixirCost)
				}
				return
			}

			if got.ElixirCost == nil {
				t.Fatal("ElixirCost = nil, want value")
			}
			if *got.ElixirCost != tt.wantValue {
				t.Errorf("ElixirCost = %d, want %d", *got.ElixirCost, tt.wantValue)
			}
		})
	}
}
