//go:build integration

package crgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	crgo "github.com/clagon/crgo"
	"github.com/clagon/crgo/model"
)

const integrationRequestTimeout = 10 * time.Second

func TestIntegrationGetCards(t *testing.T) {
	token := os.Getenv("CLASH_ROYALE_API_TOKEN")
	if token == "" {
		t.Fatal("CLASH_ROYALE_API_TOKEN must be set when running integration tests")
	}

	client, err := crgo.NewClient(token)
	if err != nil {
		t.Fatal(err)
	}

	var cards *model.Items

	t.Run("get_cards", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), integrationRequestTimeout)
		defer cancel()

		cards, err = client.GetCards(ctx, nil)
		if err != nil {
			t.Fatalf("get cards: %v", err)
		}
	})

	t.Run("cards_not_empty", func(t *testing.T) {
		if cards == nil || len(cards.Items) == 0 {
			t.Fatal("expected at least one card")
		}
	})

	t.Run("request_timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		defer cancel()
		<-ctx.Done()

		_, err := client.GetCards(ctx, nil)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("expected deadline exceeded, got %v", err)
		}
	})
}
