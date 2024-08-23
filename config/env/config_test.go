package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNumber(t *testing.T) {
	t.Run("valid integer value", func(t *testing.T) {
		err := os.Setenv("TEST_KEY", "42")
		require.NoError(t, err)
		defer os.Unsetenv("TEST_KEY")

		result, err := getNumber("TEST_KEY")

		assert.NoError(t, err)
		assert.Equal(t, 42, result)
	})

	t.Run("invalid integer value", func(t *testing.T) {
		err := os.Setenv("TEST_KEY", "not-a-number")
		require.NoError(t, err)
		defer os.Unsetenv("TEST_KEY")

		result, err := getNumber("TEST_KEY")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
		assert.Contains(t, err.Error(), "invalid TEST_KEY")
	})

	t.Run("empty environment variable", func(t *testing.T) {
		err := os.Unsetenv("TEST_KEY")
		require.NoError(t, err)

		result, err := getNumber("TEST_KEY")

		assert.Error(t, err)
		assert.Equal(t, 0, result)
		assert.Contains(t, err.Error(), "invalid TEST_KEY")
	})
}
