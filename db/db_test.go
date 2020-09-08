package db_test

import (
	"context"
	"testing"

	"github.com/go-snart/snart/test"
)

func TestDB(t *testing.T) {
	ctx := context.Background()
	d := test.DB(ctx)

	if d == nil {
		t.Error("d == nil")
	}

	t.Run("name", func(t *testing.T) {
		if d.Name != test.DBName {
			t.Errorf("d.Name == %q != test.DBName == %q", d.Name, test.DBName)
		}
	})

	t.Run("ping", func(t *testing.T) {
		ctx := context.Background()
		err := d.Ping(ctx).Err()
		if err != nil {
			t.Errorf("ping err: %w", err)
		}
	})
}
