package database

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("test_init_db", func(t *testing.T) {
		Init()
	})
}
